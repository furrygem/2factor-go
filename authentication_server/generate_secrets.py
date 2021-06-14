#!/usr/bin/env python3

from Cryptodome.IO import PEM
import secrets, sys

secrets_folder = sys.argv[1]
sys.stdout.write("Beginig secrets generation\n")
try:
    # Generating session auth secret
    key = secrets.token_bytes(64)
    # Encoding session auth secret and writing it to file
    pem = PEM.encode(key, "SESSION AUTHENTICATION SECRET")
    file = open(secrets_folder + "/session_auth.pem", "w")
    file.write(pem)
    # closing file
    file.close()
    ## DONE CREATING session authentication secret ##

    # Generating session encryption secret
    key = secrets.token_bytes(32)
    # Encoding session auth secret and writing it to file
    pem = PEM.encode(key, "SESSION ENCRYPTION SECRET")
    file = open(secrets_folder + "/session_enc.pem", "w")
    file.write(pem)
    # closing file
    file.close()
except IndexError:  # in there is no aruments is argv catching IndexError and displaying error. exit code 1
    sys.stderr.write("No output folder name provided\n")
    # exiting with error code 1
    exit(1)
except FileNotFoundError:  # if foldername provided by user is not a folder cathing FileNotFoundError and displaying error. exit code 1
    sys.stderr.write("Folder name provided but this folder does not exist\n")
    sys.stdout.write("Create folder before executing this script\n")
    # exiting with error code 1
    exit(1)
except NotADirectoryError:  # if there is a file with provided name but it is not a folder cathing NotADirectoryError and displaying error. exit code 1
    sys.stderr.write("File with provided filename exists but it is not a folder\n")
    sys.stdout.write("Make sure to provide path to a FOLDER.\n")
    # exiting with error code 1
    exit(1)
except Exception as e:  # catching other errors and displaying them. exit code 1
    sys.stderr.write("Unknown error catched\n")
    sys.stderr.write(e)
    sys.stdout.write("\n")
    # exiting with error code 1
    exit(1)
## DONE ##
sys.stdout.write("Done\n")
