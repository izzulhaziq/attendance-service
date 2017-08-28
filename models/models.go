package models

// TODO: Probably it's worth breaking this file into separate files once
// model starts growing.

// BaseModel define the base for all model, typically providing
// the unique ID of the instance of a model.
type BaseModel struct {
	ID string `json:"id"`
}

// Student represents the student model.
type Student struct {
	BaseModel
	Name      string `json:"name,omitempty"`
	StudentID string `json:"studentId,omitempty"`
	Batch     string `json:"batch,omitempty"` // e.g: class of 2003
}

// Lecturer represents the lecturer model.
type Lecturer struct {
	BaseModel
	Name     string   `json:"name,omitempty"`
	Division string   `json:"division,omitempty"` // e.g: Business school
	Courses  []Course `json:"courses,omitempty"`
}

// Course represents a subject or class that a lecturer is teaching
// and student is taking.
type Course struct {
	BaseModel
	Name string `json:"name,omitempty"`
}

// NewStudent creates a new instance of Student.
func NewStudent(name, studentID, batch string) Student {
	return Student{
		BaseModel: BaseModel{studentID},
		Name:      name,
		StudentID: studentID,
		Batch:     batch,
	}
}
