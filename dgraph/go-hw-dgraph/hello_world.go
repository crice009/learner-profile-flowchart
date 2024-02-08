/*
 * EXAMPLE USE OF DGRAPH CLIENT -- Feb 2024 -- Corey Rice
 * This is an example of how to use the dgraph client, and should not be
 * included in the project code. Rather, to serve as copy/paste inspiration.
 *
 * The 'getSchema' example shows how to interact with the schema,
 *   and this should only need done infrequently - to set up the database.
 *
 * The 'setObject' example shows how to interact with the data,
 *   and this should be done frequently - to add, update, or delete data.
 *
 */

package main

func main() {
	Example_getSchema() //returns a schema
	// Example_setObject() //returns objects
}
