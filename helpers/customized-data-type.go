package helpers

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

// UUIDType
type UUIDType uuid.UUID

// StringToUUIDType -> parse string to MYTYPE
func StringToUUID(s string) (UUIDType, error) {
	id, err := uuid.Parse(s)
	return UUIDType(id), err
}

//String -> String Representation of Binary16
func (u UUIDType) String() string {
	return uuid.UUID(u).String()
}

//GormDataType -> sets type to binary(16)
func (u UUIDType) GormDataType() string {
	return "binary(16)"
}

func (my UUIDType) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *UUIDType) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = UUIDType(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *UUIDType) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = UUIDType(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my UUIDType) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}


