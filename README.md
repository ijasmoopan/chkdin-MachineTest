# chkdin-MachineTest

This is a simple user registration and login assignment. I have done all the APIs according to the assignment. 

First of all I followed MVC architecture to do this assignment. So we can understand the source code easily.

In go there is default router. But here I used Chi router. Chi router is an open source router package like Gorilla Mux. I prefer Chi router than Gorilla Mux. 

For getting logging details, I used a logging middleware of Chi middleware.

I used PostgreSQL as database. I didn't used ORM, in Go, it is GORM. Because GORM make us very easy. But first of all we should need good knowledge in Query. So I didn't used GORM now.

The next one I have used in this assignment is an open source package named as Goose Package. Goose package is for migrating database. We can migrate database by just one up command.

The commands of goose package little big. So for making that easy I used makefile, and wrote big commands in makefile. So while calling big command will become very easy and the main thing we will get some more time to do other stuffs.

The users can register their account by just providing email and password. Then only then can login to project.

In this assignment I have done authentication. This time I used JWT authentication for authorizing the users. I have done this by using an authorizing middleware. If the user logged in, that time they will get token. They can access other pages by this token only. If they try to access any page except signup and login, they will get a message as "please login".

I used Post method for signup and login.
Then Get method is for getting user details for that particular user.
If the user wants to change the details of his account, for this I used Patch method instead of Put method. In my assignment mentioned that use Put method. But Put is for changing whole the details. But in this assignment I have done both single edit or whole edit. That means user can edit both a single details or whole details at a time by using single API.
For deleting user account I used delete method.

For the database connection I implemented the dependency injection method by using interfaces. 

In a project, the important thing is its security. In this assignment I have used some credetails. So for making those values secure, I have used environment variables. 

Thank you for reading :) .
