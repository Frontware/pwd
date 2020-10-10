# pwd
Golang libraries to manage password functions


# IsCommon

The function retruns trye if the password is in the list of 1 million most used password.
When a user signs up, you can use this function to make sure he is not providing a password that we can easily find in dicitonaries.

The function uses **Bloomfilter** to keep a low memory use.

# HashPassword

Returns a bcrypt hash of the password


# NewPassword

Returns a new password.
