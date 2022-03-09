# Keepintouch App

A simple app to keep in touch with friends and family.

## General Requirements
* Remind the user daily to get in touch with a specific person by email
  * Provide the user the contacts of the person
* Keeps track of the people you have contacted and how frequently
* Can generate a simple report of the people you keep in touch with

## Usage Requirements
* Simple app used by one user at a time
* Manage contacts of up to 1000 people

## API
* addUser(user)
* retrieveUserInfo(user_name)
* updateUser(user_name, fields_to_update)
* deleteUser(user_name)
* addusers(list_of_users)
* setNumOfDailyContact(number)
* todaysContact()
* setUpCronJob()

### Data model
* user:
  ```
  name
  phone_number (optional)
  email (optional)
  birthday (optional)
  importance_level (required)
  notes (optional)
  ```
  * Quick question: do we want to include the option for users to define random fields? No
    * Follow up to that, how do we make those fields consistent? We enforce consistency.

* list_of_users:
  * will a json file or csv file work better here?

### Data storage
* It is a simple app. We can use the following:
  * a text file
  * sqlite - let's use this. but it means we will have a rigorous data model
    * we will have some ACID properties though

### Other tasks
* How do we determine the frequency of contact?
  * do we use importance levels and map that to a certain frequency?

* Can user upgrade important people? Yes, just update user

* Can user select how many people they want to contact daily? yes

* How can user give the app quick feedback? Not in this version 

## Code Organization

* All actions have to be logged
* Expose minimal database structures 
  * The api can only say:
    * add user
    * update user
    * remove user
    * get user
* Config file to setup:
  * Num of people to contact daily
  * Email to send the people's contact
* KeepInMind APIs

## Architecture

![diagram](arch.png)

