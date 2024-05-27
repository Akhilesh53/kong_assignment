**KONG ASSIGNMENT**
------------------------
*Requirements*
------------------------
The Product Owner delivered the following story:

As a user, I can see an overview of services in my organization. Acceptance criteria include:

User can see the name, a brief description, and versions available for a given service
User can navigate to a given service from its card
User can search for a specific service
User can paginate and sort through services
The Assignment

You're responsible for the data model and API portions of this story.
Implement a Services API that can be used to implement this dashboard widget.
It should support
Returning a list of services
support filtering, sorting, pagination
Fetching a particular service
including a method for retrieving its versions

The API can be read-only. Choose a persistence mechanism that is appropriate for this feature.

This project must be written in Go.

We'll evaluate the design, implementation choices, and functionality of your project. The code should be as production ready as you can make it, even if that means reducing the total features you are able to add.

-----------------------------
*Objects Classification*
-----------------------------
You can find objects declaration at path ./pkg/models

-----------------------------
*Design Pattern*
-----------------------------
1. Here I have tried to use Singleton Design Pattern to initialise DB , initialise Logger, Env Vars.
   Reason for doing so, is to initialise all these only once, even though if same function is called again, it won't initialise a new instance.

2. Directory Structure I have tried to follow is DDD (data driven design).

-----------------------------
*End Points and Explanation*
-----------------------------

1) /services : It will fetch all the services present in database.

   Params:
    - page :  supports pagination. contains number which tells from which offset, we have to fetch the services. Default value is 1.
    - limit: limits the number of response. Default value is 10.
    - sort: sorts the result by mentioned tag. Default by ID
    - order: the order in which result gets sort. Default is asc. Possibble values are asc/ desc.


2) /services/:id  : It will fetch all the service with particular service id.

3) /service/:id/versions : It will fetch all the versions for a particular service id.

------------------------------
*Middlewares*
------------------------------

I am using two middlewares for this assignment.

1) Generate uuid : It will generate and assigned a unique id to every api request received.

2) Authorisation: It will authorise a user to allow access to hit the api or not.

-------------------------------
*Authorisation*
-------------------------------

Here I have used Basic Auth to authorise the user. 
Need to pass userid and password in basic auth section or pass a Basic Auth Token in the headers section.
If user is valid, it will authorise the access ti hit the api.

-------------------------------
*Changes to be done to make this code running on other server*
-------------------------------

1) Change the DB env vars
   - Go to ./config/config.go
   - Change the value for postgresHost, postgresPort,postgresUser,postgresPassword,postgresDB,processName

2) Create Tables and Insert data to the tables
   - Attached sql queries in ./queries folder

------------------------------
*DB Schema Design*
------------------------------

- Cardinality Relationship
   
  - Services - Version
    1 service can have M (many) version
    1 Version of service belongs to 1 service
    So, it has 1-M relationship

    In case of 1-M relationship, primary key of 1 table should be included in M table
    
Thus there will be 2 tables
Service - Contains service details
Versions - Contains version details for services.

------------------------------
*Sample Curls*
------------------------------

You can find the sample curls attached in curls folder in the directory.

