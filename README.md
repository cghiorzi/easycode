EasyCode
=

EasyCode is a simple webpage that is designed to highlight the ease with which changes can be pushed to production using Cloud Foundry. 
The page will display a greeting as well as a Google map of the location specified in `getName()` function in the index.html file (located in the `src/main/resources/templates` directory). The code 
will accept any location that Google Maps accepts, such as:

*	City and State. e.g. "Philadelphia, PA"
*	Airport codes. e.g. "BOS"
*	Addresses. e.g. "176 South Street Hopkinton, MA 01748"
*	Landmarks. e.g. "Grand Canyon"


To Install:
-
* Install [Java SDK](http://www.oracle.com/technetwork/java/javase/downloads/index.html)
* Install [Apache Maven](https://maven.apache.org/download.cgi)
* Git clone this repo
* Run `mvn package` in the top level easycode directory
* Edit the Cloud Foundry `manifest.yml` file and change the `name` and `host` parameters to something unique
* `cf push` to send to Cloud Foundry

To Update:
-
* Edit the `getName()` function in the `src/main/resources/templates/index.html` file 
* Run `mvn package` in the top level easycode directory to recompile the Java jar
* `cf push` to send the updated code to Cloud Foundry