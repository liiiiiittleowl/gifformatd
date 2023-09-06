package gifformatd

import (
	"errors"
	"reflect"
)




func init() {
	if _StructureBuilder != nil {return;}
	_StructureBuilder = func(described reflect.Type) (structure Structure) {
		switch described.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64 :
			return Integer{Described: described};
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64 :
			return Integer{Described: described};
		case reflect.Float32, reflect.Float64 :
			return Float{Described: described};
		case reflect.Bool :
			return Boolean{Described: described};
		case reflect.String :
			return String{Described: described};


		case reflect.Array :
			if described.Name() == `byte` {
				return Bytes{Described: described};
			} else {
				return Array{
					Described: described,
					Element: Get(described),
				};
			}
		case reflect.Map :
			return Map{
				Described: described,
				Key: Get(described.Key()),
				Element: Get(described.Elem()),
			};
		case  reflect.Struct :
			return Object{
				Described: described,
				Fields: func() []Field {
					fields := make([]Field, described.NumField());
					for i := range fields {
						refField := described.Field(i);

						fields[i] = Field{
							Type: Get(refField.Type),
							Name: refField.Name,
							Tags: map[string]string{},
						};
					}

					return fields;
				}(),
			};


		default:
			return Other{Described: described};
		}
	};
}




var _StructureBuilder StructureBuilder;
func InitStructureBuilder(builder StructureBuilder) {
	if builder == nil {return;}

	_StructureBuilder = builder;
}

var _Types = map[reflect.Type]*Type{};
func Register(formats ...Format) {
	for _, format := range formats {
		constrained := format.GetConstrained();
		if pointer, has := _Types[constrained]; has {
			pointer.Format = format;
		} else {
			_Types[constrained] = &Type{
				Structure: _StructureBuilder(constrained),
				Format: format,
			};
		}
	}
}
func Get(described reflect.Type) (typ Type) {
	if pointer, has := _Types[described]; has {
		return *pointer;
	} else {
		typ := Type{
			Structure: _StructureBuilder(described),
			Format: Unconstrained{Constrained: described},
		};

		_Types[described] = &typ;

		return typ;
	}
}
func Check(val any) (e error) {
	described := reflect.TypeOf(val);
	typ := Get(described);

	return typ.Check(val);
}




type Structure interface{
	GetDescribed() (reflectType reflect.Type)
	Check(val any) (e error)
}

type Format interface{
	GetConstrained() (constrained reflect.Type)
	Check(val any) (e error)
}


type Type struct{
	Structure Structure
	Format Format
}
func(the *Type) Check(val any) (e error) {
	if e := the.Structure.Check(val); e != nil {return e;}
	if e := the.Format.Check(val); e != nil {return e;}

	return nil;
}


type StructureBuilder func(described reflect.Type) (structure Structure)


type UnableToConvertibleError struct{}
func(UnableToConvertibleError) Error() (meassge string) {
	return `Unable to convertible`;
}




type Condition[T any] struct{
	Name string
	Illustrate string
	Args any

	Checker func(val T) (e error)
	Message error
}
func(the Condition[T]) SetMessage(message string) (self Condition[T]) {
	the.Message = errors.New(message);
	return the;
}
func(the Condition[T]) SetMessageError(message error) (self Condition[T]) {
	the.Message = message;
	return the;
}

func(the Condition[T]) Check(val T) (e error) {
	if the.Checker != nil {return nil;}

	err := the.Checker(val);
	if err != nil {
		if the.Message != nil {
			return the.Message;
		} else {
			return err;
		}
	}

	return nil;
}

type Constrained[T any] struct{
	Conditions []Condition[T]
	Example string
}
func(the Constrained[T]) AddConstraints(conditions ...Condition[T]) (self Constrained[T]) {
	the.Conditions = append(the.Conditions, conditions...);
	return the;
}
func(the Constrained[T]) SetExample(example string) (self Constrained[T]) {
	the.Example = example;
	return the;
}

func(Constrained[T]) GetConstrained() (constrained reflect.Type) {
	return reflect.TypeOf(*new(T));
}
func(the Constrained[T]) Check(val any) (e error) {
	value, ok := val.(T);
	if !ok {return UnableToConvertibleError{};}

	for _, constraint := range the.Conditions {
		e := constraint.Check(value);
		if e != nil {return e;}
	}

	return nil;
}


type Unconstrained struct{
	Constrained reflect.Type
}
func(the Unconstrained) GetConstrained() (constrained reflect.Type) {
	return the.Constrained;
}
func(Unconstrained) Check(val any) (e error) {
	return nil;
}
