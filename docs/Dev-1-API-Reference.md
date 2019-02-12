MikaPost API Reference
======

## Developers Notes

1. To help make the next few API endpoints easy to type, save your token to the console.

    ```bash
    MIKAPOST_API_TOKEN='YOUR_TOKEN'
    ```

2. You will notice ``http`` used in the sample calls through this document, this is the ``Python`` command line application called ``HTTPie``. Download the command line application by following [these instructions](https://httpie.org/).


3. If you are going to make any contributions, please make sure your edits follow the [API documentation standard](https://gist.github.com/iros/3426278) for this document; in addition, please read [Googles API Design Guide](https://cloud.google.com/apis/design/) for further consideration.


## Get API Version
Returns the version information of MikaPost. This is a useful endpoint to call when you are setting up your project and you want to confirm you are able to communicate with the web-service.


* **URL**

  ``/api/v1/public/version``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "Service": "v0.1",
        "API: 'v1"
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/public/version
  ```


## Register
Submit registration details into our system to automatically create a *user* account. System return the *user* details and authentication *token*.

Created *user* accounts are automatically granted access to the system even though these accounts have not had their email verified. The system sends a verification email after creation and if the *user* does not verify in the allotted timespan, their account gets locked.

It's important to note that emails must be unique and passwords strong or else validation errors get returned.

* **URL**

  ``/api/v1/public/register``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * email
  * password
  * first_name
  * last_name


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkyOTY1MDAsInVzZXJfaWQiOjF9.QN9dyWL2dlxKgkm0xbQAmnaI6_4amHcSfqUGQ6pZbxM",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Email is not unique. Please enter another email.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/public/register \
    email=bart@mikasoftware.com \
    password=YOUR_PASSWORD \
    first_name=Bart \
    last_name=Mika
  ```


## Login
Returns the *user profile* and authentication *token* upon successful login in.

* **URL**

  ``/api/v1/public/login``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * email
  * password


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkyOTg1MDYsInVzZXJfaWQiOjF9.HrwHvfL4-1pMe7EcXEzlsxciFgK0xf2uC8BV1kfLT_c",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Email or password is incorrect.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/public/login \
    email=bart@mikasoftware.com \
    password=YOUR_PASSWORD
  ```


## Get Profile
The API endpoint used to get the *user profile details*. Only the *profile* of the
  *authenticated user* is returned.

* **URL**

  ``/api/v1/profile``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "user_id": 1
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/profile Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Create Box
Creates a *box* by an *authenticated user* in our system. When a *box* is created, you may created *things objects inside it and the *box* model is used for aggregation. Configuring the ``status`` field allows the *box* to be public or private.

* **URL**

  ``/api/v1/boxes``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * name
  * short_description
  * long_description
  * status
      - 1 = Active (Private)
      - 2 = Active (Public)
      - 3 = Archived
      - 4 = Deleted


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "id": 1,
        "name": "Mika Family",
        "short_description": "The sensors used in the    Mika family home.",
        "status": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Please fill in the name.",
        "status": "Invalid request."
    }
    ```

  OR

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Please select a valid status option.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/boxes \
    Authorization:"Bearer $MIKAPOST_API_TOKEN" \
    name="My Home" \
    short_description="The sensors in our home." \
    status=1
  ```


## List Boxes
Returns paginated list of all the *boxes*.

* **URL**

  ``/api/v1/boxes``


* **Method**

  ``GET``


* **URL Params**

  * page
  * paginate_by
      - Limit is 100


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "count": 1,
        "details": [
            {
                "id": 2,
                "name": "My Home",
                "short_description": "The sensors in our home."
            }
        ],
        "next": "",
        "page": 0,
        "pages": 1,
        "previous": ""
    }
    ```


* **Error Response**

  * **Code:** 404
  * **Content:**

    ```json
    {
        "status": "Resource not found."
    }
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/boxes page==0 paginate_by==25 Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Retrieve Box
Returns the *box* details.

* **URL**

  ``/api/v1/box/<box_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "id": 1,
        "name": "My House",
        "short_description": "The sensors used in the our home.",
        "status": 1
    }
    ```


* **Error Response**

  * **Code:** 403
  * **Content:**

    ```json
    {
        "status": "Forbidden."
    }
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/box/1 Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Update Box

**TODO: IMPLEMENT**



## Create Thing
Creates an *thing* by an *authenticated user* in our system.

* **URL**

  ``/api/v1/things``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * name
  * description


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "City": "",
        "Country": "",
        "IsAddressVisible": false,
        "Postal": "",
        "Province": "",
        "ShareKey": "",
        "StreetAddress": "",
        "StreetAddressExtra": "",
        "box_id": 2,
        "id": 4,
        "long_description": "Measures relative humidity from 0 to 100% and temperature from - to +. Polls at an interval of every minute.",
        "name": "Bedroom Humidity Sensor",
        "short_description": "Humidity Phidget sensor (HUM1000_0) located on the second floor and inside master bedroom.",
        "status": 1,
        "unit_of_measure": "%",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Please fill in the name.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/things \
    Authorization:"Bearer $MIKAPOST_API_TOKEN" \
    box_id=1 \
    name="Bedroom Humidity Sensor" \
    short_description="Humidity Phidget sensor (HUM1000_0) located on the second floor and inside master bedroom." \
    long_description="Measures relative humidity from 0 to 100% and temperature from -40°C to +85°C. Polls at an \
    interval of every minute." \
    unit_of_measure="%" \
    status=1
  ```


## List Things
Returns paginated list of all the *things*.

* **URL**

  ``/api/v1/things``


* **Method**

  ``GET``


* **URL Params**

  * page
  * paginate_by
      - Limit is 100


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "count": 1,
        "details": [
            {
                "id": 1,
                "name": "Bedroom Humidity Sensor",
                "short_description": "Humidity Phidget sensor (HUM1000_0) located on the second floor and inside master bedroom.",
                "unit_of_measure": "%"
            },
        ],
        "next": "",
        "page": 0,
        "pages": 1,
        "previous": ""
    }
    ```


* **Error Response**

  * **Code:** 404
  * **Content:**

    ```json
    {
        "status": "Resource not found."
    }
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/things page==0 paginate_by==25 Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```

  OR

  ```bash
  $ http get 127.0.0.1:8080/api/v1/things Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Retrieve Thing
Returns the *thing* details.

* **URL**

  ``/api/v1/thing/<thing_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "City": "",
        "Country": "",
        "IsAddressVisible": false,
        "Postal": "",
        "Province": "",
        "ShareKey": "",
        "StreetAddress": "",
        "StreetAddressExtra": "",
        "box_id": 2,
        "id": 4,
        "long_description": "Measures relative humidity from 0 to 100% and temperature from - to +. Polls at an interval of every minute.",
        "name": "Bedroom Humidity Sensor",
        "short_description": "Humidity Phidget sensor (HUM1000_0) located on the second floor and inside master bedroom.",
        "status": 1,
        "unit_of_measure": "%",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 403
  * **Content:**

    ```json
    {
        "status": "Forbidden."
    }
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/thing/1 Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Update Thing

**TODO: IMPLEMENT**


## Create Time-Series Datum
Creates a *time-series datum by an *authenticated user* in our system.

* **URL**

  ``/api/v1/data``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * value
  * timestamp
  * thing_id


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    $ http post 127.0.0.1:8080/api/v1/data \
      Authorization:"Bearer $MIKAPOST_API_TOKEN" \
      value=123 \
      timestamp="2019-02-01T00:00:00-05:00" \
      thing_id=1
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Please select `thing_id` that you have permission for.",
        "status": "Invalid request."
    }
    ```

  OR

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Please fill in the `thing_id`.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/data \
    Authorization:"Bearer $MIKAPOST_API_TOKEN" \
    value=123 \
    timestamp="2019-02-01T00:00:00-05:00" \
    thing_id=1
  ```


## List (Thing's) Time-Series Data
Returns paginated list of all the *time-series data* belonging to the *Thing* object. Will return ``403 Forbidden`` error if you do not have permission.

* **URL**

  ``/api/v1/thing/<thing_id>/data``


* **Method**

  ``GET``


* **URL Params**

  * page


* **Data Params**

  * thing_id


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
    "count": 5,
        "details": [
            {
                "id": 2,
                "timestamp": "0000-12-31T18:42:28-05:17",
                "value": 123
            },
            {
                "id": 3,
                "timestamp": "2018-12-31T19:00:00-05:00",
                "value": 123
            },
            {
                "id": 4,
                "timestamp": "2018-12-31T19:00:00-05:00",
                "value": 123
            },
            {
                "id": 5,
                "timestamp": "2018-12-31T19:00:00-05:00",
                "value": 123
            },
            {
                "id": 6,
                "timestamp": "2018-12-31T19:00:00-05:00",
                "value": 123
            }
        ],
        "next": "",
        "page": 0,
        "pages": 1,
        "previous": ""
    }
    ```


* **Error Response**

  * **Code:** 403
  * **Content:**

    ```json
    {
        "status": "Forbidden."
    }
    ```

  OR

  * **Code:** 404
  * **Content:**

    ```json
    {
        "status": "Resource not found."
    }
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/thing/2/data page==0 Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Retrieve Time-Series Datum
Returns the *time-series datum* details.

* **URL**

  ``/api/v1/datum/<datum_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  * datum_id


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "id": 6,
        "thing_id": 2,
        "timestamp": "2019-02-01T00:00:00-05:00",
        "value": 123
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "status": "Forbidden."
    }
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/datum/1 Authorization:"Bearer $MIKAPOST_API_TOKEN"
  ```


## Update Thing

**TODO: IMPLEMENT**
