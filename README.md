# Resources : 
1. beverages
   1. name
   2. ingredientsList
2. ingredients
   1. name
   2. available-amount


# Services :
1. beverage manager/service
   1. add new beverage
2. ingredients manager/service
   1. getIngridient
   2. update ingridient
   3. add new ingridient
   4. isAvaliable for checking is ingridient is available.
3. Dispenser service (controller)
   1. make_beverage
4. beverage reposiotory
   1. Create
5. ingredient respository
   1. add
   2. update
   3. read
6. ingridients controller
   1. AddNewIngridientAPI
   2. Re-fillIngridientAPI
   3. IsAvailable
7. beverage controller
   1. AddNewBeverageAPI
8. Dispenser controller
   1. MakeBeverageAPI
9.  transaction lock manager

# Other Aspects :
1. Logging
2. HttpServerMux
3. Monitoring
4. Security
5. Exception Handling
6. Utils
7. Config Management
8. Dependency Management

# oother related project :
1. Client OR user of machine

# Processes :
1. CRUD bevrages
   1. Create : 
      1. add entry in begrage repository
      2. return success
      3. return error if such repo exists already
   2. Read :
      1. Get beverage based on name of bevrage
   3. Update : **Not asked.**
   4. Delete : **Not asked.**

2. CRUD ingredients
   1. Create :
      1. call new ingridient api from ingridient manager
      2. _add_ entry in ingredients repo.
      3. return error if adding to repo says duplicate entry.
      4. return success
   2. read : 
      1. get ingredients data using getIngridient from ingridient manager
   3. update : using update ingridient from ingridient manager
      1. decrement by X : 
         1. take lock on ingredient 
         2. update by -X units
         3. update availibility for an ingridient.
      2. increment by X :
         1. take lock on ingredient 
         2. update by -X units
         3. update availibility for an ingridient.
   4. delete : **Not asked.**
3. Get beverage
   - req. contain the outlet no., beverage type
   1. acquire lock on the outlet using the no. (and decrement the count outlet available for serving **not needeed totally**)
   2. Add entry in the orders table as "Requested" **Not asked.**
   3. Get beverage data from beverage repo.(get the ingridients req. and quanitity of each), 
   4. if couldn;t find, update to "bad-order" **Not asked.**
   5. return error.
   6. update the ingredients to lock for that many ingridients, 
      1. if any of it fails, fail the order "Unserviceabile", else udpate order to "preparing". **Not asked.**
   7. return failure if can;t be serviced.
   8. return statement : "XYZ is prepared" or somethings, check the ques.
4. add ingredients
   1. acquire lock on ingridient.
   2. update to +x OR -x, check if the quantity if postive if not return error.
5. check which ingredients are finished, track finished ingredients
   1. READ the ingridients
6. start machine
   1. initiantiate dependencies
   2. add seed data form file
7.  stop machine
    1.  graceful shutdown ie. serve the one last beverage from machine from each outlet.