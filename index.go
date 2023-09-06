package gifformatd

import (
	"errors"
	"reflect"
)




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
