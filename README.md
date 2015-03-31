# golastic

[![Build Status](https://drone.io/github.com/cone/golastic/status.png)](https://drone.io/github.com/cone/golastic/latest)
[![views](https://sourcegraph.com/api/repos/github.com/cone/golastic/.counters/views.svg)](https://sourcegraph.com/github.com/cone/golastic)

Go library to consume the elastic search API

##Getting started

###Installing the project

Simply execute the following command

    go get github.com/cone/golastic.git

###Dependencies

Golastic uses go-uuid to generate unique ids fo the indexed documents
if none was specified, so you must get the go-uuid project with the
following command

    go get code.google.com/p/go-uuid/uuid

###Use

First you must specify the url of the server you want to connect to.
Then you can start querying for or indexing documents

    server, _ := New("http://localhost:9200")

    resultItem, err := golastic.From("test", "product").Find("1")

##Operations

###From

It allows you to specify the Index and Type in a single step

###Index

It allows you to specify the Index. For example, you can use it to
delete an index

	  _, err = golastic.Index("test").DeleteDoc("")

###Find

It allows you to find a document by its id

    resultItem, err := golastic.From("test", "product").Find("1")

###Exec

It allows you to find multiple results using a query. You can specify
a method like "post" (default if left blank) or "delete". The last
in case you want to delete the matching documents

    
	result, err := golastic.From("test", "product").Exec("post", Query("match_all"))

###IndexDoc

It allows you to index a document. You can specify a custom id or leave it blank to let Golastic generate one for you


    _, err = golastic.From("test", "words").IndexDoc("1", TestProduct{1, "Potatoe"})

###UpdateDoc

It allows you to update a document sereahcing for it using its id

	  _, err = golastic.From("test", "words").UpdateDoc("1", TestProduct{2, "Apple"})

###DeleteDoc

It allows you to delete a document with the specified id

	  _, err = golastic.From("test", "words").DeleteDoc("1")

###Bulk

It allows you to do bulk operations

    b := Bulk().
        Index("1", TestProduct{1, "Prod1"}).
        Index("2", TestProduct{2, "Prod2"}).
        Update("3", TestProduct{3, "Prod3"})


    errs := golastic.From("test", "words").Bulk(b)

###Scan

It lets you map the resulting document to a struct. It can only be used
on results

	  resultItem, _ := golastic.From("test", "product").Find("1")

    product := TestProduct{}

    resultItem.Scan(&product)

You can also pass a slice for query operations

	  result, err := golastic.From("test", "product").Exec("", Query("match_all"))

    products := []TestProduct{}
    //or also []*TestProduct{}

    result.Scan(&products)

##Querying

The queries are executed using the ``Exec`` function. Aside the method
(e.g. "post" or "delete" or "" to default to "post") you should specify
the query as a map or using the Item helper

    //passing a map
    q = Query("match").Set(
		  map[string]map[string]string{
        "message": map[string]string{
          "query": "this is a test",
        },
      },
	  )
    
    //Results in:
    /*
      {  
        "query":{  
          "match":{  
            "message":{  
              "query":"this is a test"
            }
          }
        }
      }
    */

Or you can use the ``Item`` helper to achieve the same result as above

    q = Query("match").Set(
		  Item().Put("message", Item().Put("query", "this is a test")),
	  )

More helpers for the queries will be added in the future.

## Contributing

1. Create your feature branch (`git checkout -b feature/my-new-feature`)
2. Commit your changes (`git commit -am 'Add some feature'`)
3. Push to the branch (`git push origin feature/my-new-feature`)
4. Create a new Pull Request

## About the Author

Carlos Gutierrez (cone) is a sofware developer based in Colima, Mexico.
He is currently working for an awsome company called [Crowd Interactive](http://www.crowdint.com) a company specialized in building and growing online retail stores.
