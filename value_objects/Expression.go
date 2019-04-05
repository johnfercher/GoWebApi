package value_objects

type Expression struct {
	ActualExpression string
	Resolution       []string
}

func (e Expression) UpdateExpression(newExpression string) {
	e.ActualExpression = newExpression
	e.Resolution = append(e.Resolution, newExpression)
}
