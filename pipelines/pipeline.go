package pipelines

// ID of a pipeline.
type ID string

// Pipeline definition, this is kept as simple interface because there's constant changes.
// It's your job to do map to the struct with fields you need specifically.
type Pipeline map[string]interface{}
