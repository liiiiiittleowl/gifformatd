package gifformatd

import "reflect"




type Integer struct{
	Described reflect.Type
}
func(the Integer) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Integer) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}
	return nil;
}


type Float struct{
	Described reflect.Type
}
func(the Float) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Float) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}
	return nil;
}


type Boolean struct{
	Described reflect.Type
}
func(the Boolean) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Boolean) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}
	return nil;
}


type String struct{
	Described reflect.Type
}
func(the String) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the String) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}
	return nil;
}


type Bytes struct{
	Described reflect.Type
}
func(the Bytes) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Bytes) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}
	return nil;
}




type Array struct{
	Described reflect.Type

	Element Type
}
func(the Array) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Array) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}

	value := reflect.ValueOf(val);
	for i, len := 0, value.Len(); i< len; i++ {
		item := value.Index(i);
		err := the.Element.Check(item.Interface());
		if err != nil {return err;}
	}

	return nil;
}


type Map struct{
	Described reflect.Type

	Key Type
	Element Type
}
func(the Map) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Map) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}

	value := reflect.ValueOf(val).MapRange();
	for value.Next() {
		err := the.Key.Check(value.Key().Interface());
		if err != nil {return err;}

		err = the.Element.Check(value.Value().Interface());
		if err != nil {return err;}
	}

	return nil;
}


type Field struct{
	Type Type
	Name string
	Tags map[string]string
}
type Object struct{
	Described reflect.Type

	Fields []Field
}
func(the Object) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Object) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}

	value := reflect.ValueOf(val);
	for i, len := 0, value.NumField(); i < len; i++ {
		field := value.Field(i);
		err := the.Fields[i].Type.Check(field.Interface());
		if err != nil {return err;}
	}

	return nil;
}




type Other struct{
	Described reflect.Type
}
func(the Other) GetDescribed() (reflectType reflect.Type) {
	return the.Described;
}
func(the Other) Check(val any) (e error) {
	if !reflect.TypeOf(val).ConvertibleTo(the.Described) {return UnableToConvertibleError{};}
	return nil;
}
