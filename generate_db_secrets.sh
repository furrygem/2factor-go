#!/bin/bash

DB_SECRETS_FOLDER=database/secrets # DATABASE SECRET FOLDER
DB_SECRETS_KEY_FILE=db.key
DB_SECRETS_CERT_FILE=db.crt
DB_SECRETS_PASSWORD_FILE=dbpassword.txt

OPENSSL_EC_CURVE="prime256v1"
OPENSSL_EC_PARAMS="-newkey ec -pkeyopt ec_paramgen_curve:$OPENSSL_EC_CURVE"

echo -n "Enter password for database: "
read -s password
echo

if [[ ! -d "$DB_SECRETS_FOLDER" ]]; then # checking if directory database/secrets NOT exists
    if [[ -e "$DB_SECRETS_FOLDER" ]]; then # if there is any file named same as target directory exiting with error code 1
        echo "There is a file named $DB_SECRETS_FOLDER"
        echo "Unable to proceed"
        exit 1
    else # if any file named same as DB_SECRETS_FOLDER does not exist
        # Asking user should we create new database
        echo "There is no folder $DB_SECRETS_FOLDER"
        read -p "Create folder $DB_SECRETS_FOLDER? (y/n) " -n 1
        echo
        if [[ $REPLY =~ ^[Yy]$ ]];then # Checking reply with regex for Y or y (yes answer)
            # if yes creating folder with path and 
            echo "Creating folder with path $DB_SECRETS_FOLDER"
            mkdir -p $DB_SECRETS_FOLDER
        else # when user not answering y or Y assuming it as No and exiting script
            echo "Exiting"
            exit 0
        fi
    fi

fi

if [[ -e "$DB_SECRETS_FOLDER/$DB_SECRETS_PASSWORD_FILE" ]]; then # if file containing password for db already exists
    # Qustioning user should we replace it or leave as it is
    echo "File $DB_SECRETS_FOLDER/$DB_SECRETS_PASSWORD_FILE exists. "
    read -p "Overwrite $DB_SECRETS_PASSWORD_FILE? (y/n) " -n 1
    echo
    if [[ $REPLY =~ ^[Yy]$ ]];then # Checking reply with regex for Y or y (yes answer)
        echo "Writing new password $DB_SECRETS_PASSWORD_FILE"
        echo $password > $DB_SECRETS_FOLDER/$DB_SECRETS_PASSWORD_FILE
    else
        echo "File will not be overwritten"
    fi
else
    # if file containing password for db do not exist creating one and writing there password provided by user 
    echo $password > $DB_SECRETS_FOLDER/$DB_SECRETS_PASSWORD_FILE
fi



# Checking if DB_SECRETS_KEY_FILE OR DB_SECRETS_CERT_FILE exist
if [[ -e "$DB_SECRETS_FOLDER/$DB_SECRETS_KEY_FILE" ]] || [[ -e "$DB_SECRETS_FOLDER/$DB_SECRETS_CERT_FILE" ]]; then # if file containing key for db already exists
    # Qustioning user should we replace it or leave as it is
    echo "File $DB_SECRETS_FOLDER/$DB_SECRETS_KEY_FILE exists. Replace CERT/KEY pair? (y/n) "
    read -p "Overwrite $DB_SECRETS_KEY_FILE and $DB_SECRETS_CERT_FILE? (y/n) " -n 1
    echo
    if [[ $REPLY =~ ^[Yy]$ ]];then # Checking reply with regex for Y or y (yes answer)
        # if yes setting gen_new_openssl to true
        gen_new_openssl=true
    else
        # else setting gen_new_openssl to false
        gen_new_openssl=false
    fi
else
    # if no certificate/key file exists just generating new ones by setting gen_new_openssl to true
    gen_new_openssl=true
fi


if [ "$gen_new_openssl" = true ]; then #if replace_cert is true
    # setting openssl certificate output parameters and calling openssl
    openssl_out_parameter="-out $DB_SECRETS_FOLDER/$DB_SECRETS_CERT_FILE"
    openssl_keyout_parameter="-keyout $DB_SECRETS_FOLDER/$DB_SECRETS_KEY_FILE"
    # calling openssl and passing out and keyout parameters 
    openssl req -new $OPENSSL_EC_PARAMS -x509 -nodes $openssl_keyout_parameter $openssl_out_parameter
    echo "New openssl data generated!"
else
    echo "Not replacing openssl key/cert pair"
    exit 0
fi
