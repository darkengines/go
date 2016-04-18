package service

func Process(context interface{}) error {
	adapter, _ := GetQueryAdapter(context)
	err := adapter.Handle(context)
	if (err != nil) {
		return err
	}
	return nil
}