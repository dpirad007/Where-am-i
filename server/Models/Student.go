package Models

type Student struct {
    ID string                   `bson:"_id"`
    Name string                 `bson:"Name"`
    RollNo string               `bson:"RollNo"`
    PhoneNo []string            `bson:"PhoneNo"`
    Email []string              `bson:"Email"`
    ParentPhoneNo []string      `bson:"ParentPhoneNo"`
}

func GetStudent(rollNo string) {

}
