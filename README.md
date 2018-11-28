# AddressBook.GO

## Overview
This repository contains **Address Book** application for GoLang that shows design & coding practices followed by **[Differenz System](http://www.differenzsystem.com/)**.

The app does the following:
1. **Login:** User can login via email/password. 
2. **Home:** It will list all the save contacts, having the option to add a new contact on the top right.
3. **Create new contact:** User can add a new contact to his address book by filling details here.

## Pre-requisites
- [Visual Studio code](https://code.visualstudio.com/)
- [Git](https://git-scm.com/downloads )
- [Go](https://golang.org/doc/install)



## Getting Started
1. [Install Visual Studio code](https://code.visualstudio.com/)
2. Clone this sample repository
3. download and Install go  
4. Enter command "go run filename.go".


## Key Tools & Technologies
- **Database:** MYSQL(sequelize)
- **Authentication:** login
- **API/Service calls:** fetch API
- **IDE:**  VSCode




## API
###
Registration:
http://192.168.1.142:8800/Registration

Method :POST

-req:
```
  {
	"email":"dev@gmail.com",
	"password":"dev"
	}
 ```
-res:
```
{
    "Email": "dev@gmail.com",
    "UserID": "128"
}
```

###
login:
http://192.168.1.142:8800/login

Method :POST

-req:
```
  {
	"email":"dev@gmail.com",
	"password":"dev"
  }
```
-res:
```
valid request:
    {
    "Email": "dev@gmail.com",
    "UserID": "126"
    }
invalid request:
    {
    "msg": "invalid email or password"
    }
```

###
Display addressbook:
http://192.168.1.142:8080/ShowAddress/{userid}

Method :GET

-res:
```
    [
    {
        "address_id": "77",
        "contact_number": "789456123767",
        "email": "om@12ddd",
        "is_active": "0",
        "user_id": "105"
    },
    {
        "address_id": "78",
        "contact_number": "789456123767",
        "email": "om@12ddd",
        "is_active": "0",
        "user_id": "105"
    },
    {
        "address_id": "79",
        "contact_number": "789456123767",
        "email": "om@1",
        "is_active": "0",
        "user_id": "105"
    },
    {
        "address_id": "81",
        "contact_number": "95655566",
        "email": "ass@",
        "is_active": "0",
        "user_id": "105"
    },
    {
        "address_id": "82",
        "contact_number": "95655566",
        "email": "as@dd",
        "is_active": "0",
        "user_id": "105"
    }
]
```
###
Add Address
http://localhost:8800/AddAddress

Method :POST

-req:
```
  {
	"name" : "Sam",
	"Email":"sam@gmail.com",
	"ContactNumber": "7788445566",
	"isActive" : 0,
	"userid": 105
  }

```
-res:
```
{
    "address_id": "83",
    "contact_number": "7788445566",
    "email": "sam@gmail.com",
    "is_active": "0",
    "name": "Sam",
    "user_id": "105"
}
```
###
Update Address
http://192.168.1.142:8800/UpdateAddress/{addressid}

Method :POST

-req:
```
   {
	"name" : "Sam patel",
	"Email":"sam@gmail.com",
	"ContactNumber": "7788445566",
	"isActive" : 0,
	"userid": 105
}

```
-res:
```	{
    "address_id": "83",
    "contact_number": "7788445566",
    "email": "sam@gmail.com",
    "is_active": "0",
    "name": "Sam patel",
    "user_id": "105"
}
```
###
Delete Address
http://localhost:8800/DeleteAddress/{addressid}

Method :GET

-res
 ```
 {
    "success": "Delete Successfully."
 }
```
.

## Support
If you've found an error in this sample, please [report an issue](https://github.com/differenz-system/Addressbook.Go). You can also send your feedback and suggestions at info@differenzsystem.com

Happy coding!
