package entity

type User struct {
	ID    string
	Name  string
	Email string
}

// database, schema, table
// schema: nhosm cac table, (postgrsql: public schema)

// database spady (user_schema(user))
// điểm dỡ vẫn khó scale, điểm ok là đã cô lập được các database

// order
// order_table

// porduct
// product_table
