package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tasks", func() {
	Title("Task Service")
	Description("Service for manipulate task data")
	Server("tasks", func() {
		Description("Hosts the Task Service.")

		// List the services hosted by this server.
		Services("tasks", "openapi")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://0.0.0.0:8001")
			URI("grpc://0.0.0.0:8081")
		})
	})
})

// StoredUser is the result of user data
var StoredUser = ResultType("application/vnd.stored-user", func() {
	Description("A StoredUser describes a user retrieved by the users service.")
	Reference(User)
	TypeName("StoredUser")

	Attributes(func() {
		Field(1, "email")
		Field(2, "firstname")
		Field(3, "lastname")
		Field(4, "isactive")
		Field(5, "role")
	})

	View("default", func() {
		Attribute("email")
		Attribute("role")
		Attribute("firstname")
		Attribute("lastname")
		Attribute("isactive")
	})

	View("tiny", func() {
		Attribute("email")
		Attribute("role")
		Attribute("isactive")
	})

	Required("email", "firstname", "lastname", "role")
})

// User type
var User = Type("User", func() {
	Description("User describes a user to be stored.")
	Field(1, "email", String, "Email of the user", func() {
		Pattern(`.+@.+\..{1,6}`)
		Example("ehabterra@hotmail.com")
	})
	Field(2, "firstname", String, "First Name of the user", func() {
		MaxLength(100)
		Example("Ehab")
	})
	Field(3, "lastname", String, "Last Name of user", func() {
		MaxLength(100)
		Example("Terra")
	})
	Field(4, "role", String, "user role", func() {
		Example("admin")
		Pattern(`[a-z]+[a-z0-9]*`)
	})
	Field(5, "isactive", Boolean, "Is user active.", func() {
		Default(true)
	})
	Required("email", "firstname", "lastname", "role")
})

// StoredTask is the result of task data
var StoredTask = ResultType("application/vnd.stored-task", func() {
	Description("A StoredTask describes a task retrieved by the tasks service.")
	Reference(Task)
	TypeName("StoredTask")

	Attributes(func() {
		Field(1, "id")
		Field(2, "title")
		Field(3, "description")
		Field(4, "created_date")
		Field(5, "updated_date")
		Field(6, "due_date")
		Field(7, "status")
		Field(8, "owner")
		Field(9, "assignee")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("created_date")
		Attribute("updated_date")
		Attribute("due_date")
		Attribute("status")
		Attribute("owner", StoredUser, func() {
			View("tiny")
		})
		Attribute("assignee", StoredUser, func() {
			View("tiny")
		})
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("title")
		Attribute("assignee", StoredUser, func() {
			View("tiny")
		})
		Attribute("status")
	})

	Required("id", "title", "description", "created_date", "status", "owner")
})

// Task type
var Task = Type("Task", func() {
	Description("Task describes a task to be stored.")
	Field(1, "title", String, "Title of the task", func() {
		MaxLength(200)
		Example("New task title")
	})
	Field(2, "description", String, "Description of the task", func() {
		MaxLength(5000)
		Example("Task description")
	})
	Field(3, "created_date", String, "Created date", func() {
		Format(FormatDateTime)
	})
	Field(4, "updated_date", String, "Udated date", func() {
		Format(FormatDateTime)
	})
	Field(5, "due_date", String, "due date", func() {
		Format(FormatDateTime)
	})
	Field(6, "status", String, "Status.", func() {
		Enum("Open", "Closed", "Pending")
		Default("Open")
	})
	Field(7, "owner", StoredUser, "Owner.", func() {
	})
	Field(8, "assignee", StoredUser, "Assignee.", func() {
	})
	Required("title", "description", "created_date", "updated_date", "status")
})

// NotFound type
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a task that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("task 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing task")
	Required("message", "id")
})

var _ = Service("tasks", func() {
	Description("The tasks service performs task data.")

	HTTP(func() {
		Path("/tasks")
	})

	Method("list", func() {
		Description("List all stored tasks")
		Payload(func() {
			Field(1, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
		})
		Result(CollectionOf(StoredTask))
		HTTP(func() {
			GET("/")
			Param("view")
			Response(StatusOK)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show task by ID")
		Payload(func() {
			Field(1, "id", String, "ID of task to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StoredTask)
		Error("not_found", NotFound, "Task not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new task and return ID.")
		Payload(Task)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description("Update existing task and return ID.")
		Payload(func() {
			Field(1, "id", String, "ID of task to show")
			Field(2, "task", StoredTask)
			Required("id", "task")
		})
		Result(String)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove task from tasks data")
		Payload(func() {
			Field(1, "id", String, "ID of task to remove")
			Required("id")
		})
		Error("not_found", NotFound, "ID not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("status", func() {
		Description("change task status by id")
		Payload(func() {
			Field(1, "id", String, "ID of task")
			Field(2, "status", String, "Status.", func() {
				Enum("Open", "Closed", "Pending")
				Default("Open")
			})
			Required("id", "status")
		})
		HTTP(func() {
			PUT("/status")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

})

var _ = Service("openapi", func() {
	Meta("swagger:generate", "false")
	HTTP(func() {
		Path("/")
	})
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger/{*filepath}", "./public/swagger/")
	Files("/swagger.json", "./gen/http/openapi.json")
})
