package window

import (
	"database/sql"
	"fmt"
	database "main/modules/DataBase"
	query "main/modules/querys"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myWindow_master fyne.Window

//var myApp fyne.App
var db *sql.DB

func Window_master() {
	myApp := app.New()
	myWindow_master = myApp.NewWindow("Window Master")

	db = database.Prueba_con_bd()

	loadDataAndShow()
	myWindow_master.ShowAndRun()

	defer db.Close()

}

func loadDataAndShow() {

	data_product := query.Return_id_product(db)
	data_client := query.Return_id_client(db)

	//creacion de listas que muestran los clientes y productos
	var productWidgets []fyne.CanvasObject
	for _, row := range data_product {
		productID, productName := row[0], row[1]
		label := widget.NewLabel(fmt.Sprintf("Product ID: %s, Name: %s", productID, productName))
		productWidgets = append(productWidgets, label)
	}

	var clientWidgets []fyne.CanvasObject
	for _, row := range data_client {
		clientID, clientName, idProduct := row[0], row[1], row[2]
		label := widget.NewLabel(fmt.Sprintf("Client ID: %s, Name: %s, id producto: %s", clientID, clientName, idProduct))
		clientWidgets = append(clientWidgets, label)
	}

	/* CREACION DE LOS LABELS, CUADROS DE TEXTO O BOTONES*/

	title_new_user := widget.NewLabel("Alta de cliente")
	input_new_user_name := widget.NewEntry()
	input_new_user_name.SetPlaceHolder("Name")
	input_new_user_last_name := widget.NewEntry()
	input_new_user_last_name.SetPlaceHolder("Last Name")

	title_update_user := widget.NewLabel("Actualizar Cliente")
	input_update_old_name := widget.NewEntry()
	input_update_old_name.SetPlaceHolder("Old Name")
	input_update_new_name := widget.NewEntry()
	input_update_new_name.SetPlaceHolder("New Name")
	input_update_new_last_name := widget.NewEntry()
	input_update_new_last_name.SetPlaceHolder("New Last Name")

	title_delete_user := widget.NewLabel("Eliminar Cliente")
	input_delete_user_name := widget.NewEntry()
	input_delete_user_name.SetPlaceHolder("Name")

	title_new_product_ := widget.NewLabel("Alta de producto")
	input_new_product_name := widget.NewEntry()
	input_new_product_name.SetPlaceHolder("Product Name")
	input_new_product_description := widget.NewEntry()
	input_new_product_description.SetPlaceHolder("Description")
	/*
		title_update_product := widget.NewLabel("Actualizar Producto")
		input_update_old_product_name := widget.NewEntry()
		input_update_old_product_name.SetPlaceHolder("Old Product Name")
		input_update_new_product_name := widget.NewEntry()
		input_update_new_product_name.SetPlaceHolder("New Product Name")
	*/
	title_delete_product := widget.NewLabel("Eliminar Producto")
	input_delete_product_name := widget.NewEntry()
	input_delete_product_name.SetPlaceHolder("Product Name")

	title_product_user := widget.NewLabel("Carga de prodcuto a cliente, se necesita el nombre o id y el id del producto")
	title_new_product := widget.NewLabel("Agregar Producto a Cliente")
	input_new_product := widget.NewEntry()
	input_new_product.SetPlaceHolder("id Product")
	title_user_product := widget.NewLabel("Name the client")
	input_user_product := widget.NewEntry()
	input_user_product.SetPlaceHolder("name")

	// ACA EMPIEZA LA CREACION DEL CONTENT, QUE SERIA EL QUE ALMACENA
	//TODO LOS WIDGETS

	content := container.NewVBox(
		title_new_user,
		input_new_user_name,
		input_new_user_last_name,
		widget.NewButton("Save", func() {
			// Acción para guardar nuevo usuario
			name := input_new_user_name.Text
			last_name := input_new_user_last_name.Text

			// Agrega aquí la lógica para guardar el nuevo usuario en la base de datos
			query.Load_user(db, name, last_name)

			//aplicamos recursion para recargar la pagina a la hora de tocar el boton
			loadDataAndShow()
		}),

		title_update_user,
		input_update_old_name,
		input_update_new_name,
		input_update_new_last_name, // Nuevo campo New Last Name
		widget.NewButton("Save", func() {
			// Acción para actualizar usuario
			oldName := input_update_old_name.Text
			newName := input_update_new_name.Text
			newLastName := input_update_new_last_name.Text

			// Agrega aquí la lógica para actualizar el usuario en la base de datos
			query.Update_user(db, oldName, newName, newLastName)

			//aplicamos recursion para recargar la pagina a la hora de tocar el boton
			loadDataAndShow()
		}),

		title_delete_user,
		input_delete_user_name,
		widget.NewButton("Delete", func() {
			// Acción para eliminar usuario
			userNameToDelete := input_delete_user_name.Text

			// Agrega aquí la lógica para eliminar el usuario de la base de datos
			query.Delete_user(db, userNameToDelete)

			//aplicamos recursion para recargar la pagina a la hora de tocar el boton
			loadDataAndShow()
		}),

		title_new_product_,
		input_new_product_name,
		input_new_product_description,
		widget.NewButton("Save", func() {
			name := input_new_product_name.Text
			descripcion := input_new_product_description.Text

			query.Load_product(db, name, descripcion)

			loadDataAndShow()
		}),

		title_delete_product,
		input_delete_product_name,
		widget.NewButton("Delete", func() {
			name := input_delete_product_name.Text

			query.Delete_product(db, name)

			loadDataAndShow()
		}),

		title_product_user,
		title_user_product,
		input_user_product,
		title_new_product,
		input_new_product,
		widget.NewButton("Upload", func() {
			idProductStr := input_new_product.Text
			idProduct, err := strconv.Atoi(idProductStr)
			if err != nil {
				fmt.Print("Error")
			}

			name := input_user_product.Text

			query.New_product_user(db, idProduct, name)

			//aplicamos recursion para recargar la pagina a la hora de tocar el boton
			loadDataAndShow()
		}),

		/*widget.NewButton("Refresh", func() {
			//aplicamos recursion para recargar la pagina a la hora de tocar el boton
			loadDataAndShow()
		}),*/

		//agregamos las listas para mostrar los clientes y productos
		container.NewHBox(productWidgets...),
		container.NewHBox(clientWidgets...),
	)

	myWindow_master.SetFullScreen(true)
	myWindow_master.SetContent(content)

}
