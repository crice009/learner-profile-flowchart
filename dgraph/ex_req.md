# Example Requests
Dgraph uses it's own language for requests to the database. it is close to [GraphQL](https://graphql.org/) but not exactly the same. The examples below should show the core of how to interact with a Dgraph database.

Specifically, these are pure JSON, and designed to be handled with [Ratel](https://dgraph.io/docs/ratel/overview/), the web-interface for Dgraph databases. However, this is not how interactions will happen in an app ~ that will all be done by one of the [GraphQL clients](https://dgraph.io/docs/graphql/graphql-clients/graphql-ui/).


## First, create a Schema
! this needs redefined

https://dgraph.io/tour/schema/1/

## Run a mutation
The *Create*, *Update*, and *Delete* operations in Dgraph are called mutations.

```GraphQL 
{
   "set": [
     {
       "name":"Star Wars: Episode IV - A New Hope",
       "release_date": "1977-05-25",
       "director": {
         "name": "George Lucas",
         "dgraph.type": "Person"
       },
       "starring" : [
         {
           "name": "Luke Skywalker"
         },
         {
           "name": "Princess Leia"
         },
         {
           "name": "Han Solo"
         }
       ]
     },
     {
       "name":"Star Trek: The Motion Picture",
       "release_date": "1979-12-07"
     }
   ]
 }  
```
# Read the data
This is the first test read of the data made above.
```GraphQL
{
  movies(func: has(release_date)) {
    name
    director { name }
    starring { name }
  }
}
```
# Alter Schema
This is best done directly in Ratel interface, as described [here](https://dgraph.io/docs/dql/dql-get-started/#step-4-alter-schema).

Looking for a wa to do this without a GUI...

# Using Indexes
Once a schema has been marked to use indexes, it can be searched with 'lt'(less than) and other qualifiers. This is described [here](https://dgraph.io/docs/dql/dql-get-started/#step-5-queries-using-indexes), and can be done by command:
```GraphQL
{
  me(func: allofterms(name, "Star"), orderasc: release_date) @filter(lt(release_date, "1979")) {
    name
    release_date
    revenue
    running_time
    director {
     name
    }
    starring (orderasc: name) {
     name
    }
  }
}
```
## Next Steps
Going further can happen a number of ways, but the first steps in that progression are captured [here](https://dgraph.io/docs/dql/dql-get-started/#where-to-go-from-here).