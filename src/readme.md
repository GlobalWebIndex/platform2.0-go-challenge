Project: GlobalWebIndex Web Server
----------------------------------

Workspace Structure
-------------------
src/
  GWIWebServer/
	  assets/
      asset.go
      audience.go
      chart.go
      insight.go
      statTypes.go
    eventHandler/
      eventHandler.go
    main/
      GWIWebServer.go
    users/
      user.go


GWIWebServer.go
---------------
Main source file. Implements the endpoints of the web server and raises the corresponding events for the eventHandler (application). If an endpoint requires any parameters, the appropriate checks are performed. Also, a simple simulation is executed to fill the application's database with some random data, as a starting point.

eventHandler.go
---------------
Handles the events raised from the endpoints. Actions available: add a new user, add a new asset, display all users, display all assets, display favorites of a user, add an asset to a user's favorites, delete an asset from a user's favorites, edit an asset's description. All appropriate checks are performed and actions' result is returned as an output string to the client.

user.go
-------
Keeps user credentials and personal data. Implements functions regarding user data.

asset.go
--------
Basic asset structure which includes assets' common fields and functions.

chart.go, insight.go, audience.go
---------------------------------
Cases of asset types which each includes its own special fields and functions.

statTypes.go
------------
Declaration of variable types to which only a specific set of values can be assigned.


Endpoints description and use
-----------------------------
"http://localhost:8080/GWI-platform/show-users/":
  A list of all users is returned.

"http://localhost:8080/GWI-platform/show-assets/":
  A list of all assets is returned.

"http://localhost:8080/GWI-platform/show-favorites/?userID=XXXX":
  A list of a user's favorites is returned. Need to provide userID parameter.

"http://localhost:8080/GWI-platform/add-favorite/?userID=XXXX&assetID=YYYY":
  Adds an asset to a user's favorites. Need to provide userID and assetID parameters.

"http://localhost:8080/GWI-platform/delete-favorite/?userID=XXXX&assetID=YYYY":
  Deletes an asset from a user's favorites. Need to provide userID and assetID parameters.

"http://localhost:8080/GWI-platform/edit-asset-description/?assetID=YYYY&description=ZZZZ":
  Edits an asset's description. Need to provide assetID and description parameters.
