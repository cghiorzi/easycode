EasyCode
=====

EasyCode is a simple webpage that is designed to highlight the ease with which changes can be pushed to production using Cloud Foundry. 
The page will display a greeting as well as a Google map of the location specified in `getName()` function in the index.html file. The code 
will accept any location that Google Maps accepts, such as:
*	City and State. e.g. "Philadelphia, PA"
*	Airport codes. e.g. "BOS"
*	Addresses. e.g. "176 South Street Hopkinton, MA 01748"
*	Landmarks. e.g. "Grand Canyon"


To Install:
* Install [Java SDK](http://www.oracle.com/technetwork/java/javase/downloads/index.html)
* Install [Apache Maven](https://maven.apache.org/download.cgi)
* Git clone this repo
* Run `mvn package` in the top level easycode directory

The HTML template that contains the `getName` function is located in the `src/main/resources/templates` directory.