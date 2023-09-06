package gifformatd

import "reflect"




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
