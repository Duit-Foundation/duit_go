# Duit - drived UI tooklit.

***Duit*** is a server side UI framework for Flutter. It is used for creating widgets and server-side state management.

The framework consists of several parts:

- [Flutter package](https://github.com/lesleysin/flutter_duit)
- Go backend adapter (this repository)
- [Node JS backend adapter](https://github.com/lesleysin/duit_js)

The framework ensures that the layout model is received from the server, interacts with the backend via the Action API, and embeds custom components into the widget hierarchy processing pipeline. Duit is flexible and extensible, which allows it to create rich UI dynamically.

## Core features
- Structured mappings out of the box. UI property structures and constants.
- A simple contract for building a hierarchy of widgets.
- Ready-made widget functions
- Easily create custom actions and their dependencies

## Usage example

1. Create widget composition and build json from it
```go
func Example() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

  //Create root view from text widget. Assigning a text value and a style object
	builder.RootFrom(duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
		Data: "Example",
		Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
			FontWeight: duit_text_properties.W900,
		},
	}, "", false, nil))

  //build json
	value, err := builder.Build()

	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	return value
}
```
2. Run function and return result to client side
```go
eng.GET("/layout1", func(ctx *gin.Context) {
		view := internal.Example()
		ctx.Data(200, "application/json", view)
	})
```

## Future plans
- Widget library expansion
- Troubleshooting, updating documentation

## License
MIT


