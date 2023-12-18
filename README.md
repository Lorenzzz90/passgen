# Requirements

Dato un set di caratteri generare una password in output di 8 caratteri

## Steps

1. Check whether there are extra arguments provided to the Go run command
    1. If that's the case, check that there is only one extra argument
1. Create an array that holds all the available characters to craft the password
1. Craft a password with the available characters
1. Print the just-created password on the terminal

### Usage

1. If u want to use a custom set of characters pass them as an argument. EXAMPLE: ./passgen abcdefgh
1. If u didn't pass additional arguments the program will return an 8-character long password using the default set of characters: `abcdefgh`
