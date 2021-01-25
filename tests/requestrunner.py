


try:
    import requests
except ImportError:
    print("Missing dependencies, please run:")
    print("pip install requests")
    exit(1)

SESSION = requests.Session()

###############################################################################
# HTTP CALLS
###############################################################################


def _debug_response(debug, response):
    if debug:
        print(
            "< {status_code} {reason}".format(
                status_code=response.status_code, reason=response.reason
            )
        )

        for key, value in response.headers.items():
            print("< {key}: {value}".format(key=key, value=value))

        if response.status_code != 204:
            print("< " + str(response.content))
        else:
            print("< (no body)")


def _debug_request(debug, method, url, params, headers, body=None):
    if debug:
        print("> {method} {url} {params}".format(method=method, url=url, params=params))

        for key, value in headers.items():
            print("> {key}: {value}".format(key=key, value=value))

        if body is not None:
            print("> {body}".format(body=body))
        else:
            print("> (no body)")


def _check_status_code(response, expected_status_codes, description):
    if response.status_code not in expected_status_codes:
        print("{description} failed".format(description=description))
        sys.exit(1)


def get(url, params, headers, expected_status_codes, description, debug):
    response = SESSION.get(url, params=params, headers=headers)

    if response.status_code not in expected_status_codes:
        debug = True

    _debug_request(debug, "GET", url, params, headers)
    _debug_response(debug, response)
    _check_status_code(response, expected_status_codes, description)

    return response


def post(url, params, headers, body, expected_status_codes, description, debug):
    response = SESSION.post(url, params=params, json=body, headers=headers)

    if response.status_code not in expected_status_codes:
        debug = True

    _debug_request(debug, "POST", url, params, headers, body)
    _debug_response(debug, response)
    _check_status_code(response, expected_status_codes, description)

    return response


def put(url, params, headers, body, expected_status_codes, description, debug):
    response = SESSION.put(url, params=params, json=body, headers=headers)

    if response.status_code not in expected_status_codes:
        debug = True

    _debug_request(debug, "PUT", url, params, headers, body)
    _debug_response(debug, response)
    _check_status_code(response, expected_status_codes, description)

    return response


def delete(url, params, headers, expected_status_codes, description, debug):
    response = SESSION.delete(url, params=params, headers=headers)

    if response.status_code not in expected_status_codes:
        debug = True

    _debug_request(debug, "DELETE", url, params, headers)
    _debug_response(debug, response)
    _check_status_code(response, expected_status_codes, description)

    return response

def get_args():
    parser = argparse.ArgumentParser(
        description="A CLI for managing accounts in Uno & Umbrella"
    )
    subparsers = parser.add_subparsers()

    # Create Command
    create_parser = make_subparser(
        subparsers,
        "create",
        create_command,
        help=("Create a User Account" "Account."),
    )
    create_parser.add_argument(
        "--email",
        help=(
            "The email address to use. If not specified, an email address "
            "will be auto-generated."
        ),
    )
    create_parser.add_argument(
        "--username",
        help=(
            "The username to use. If not specified, a username will be "
            "auto-generated."
        ),
    )
    create_parser.add_argument(
        "--password",
        help=(
            "The password to use. If not specified, a password will be "
            "auto-generated."
        ),
    )

    # Get Command
    get_parser = make_subparser(
        subparsers,
        "get",
        get_command,
        help=(
            "Fetch an account, specified by UUID, Email, Username"
        ),
    )
    get_parser.add_argument("identifier", type=str)

    # Delete Command
    delete_parser = make_subparser(
        subparsers,
        "delete",
        delete_command,
        help=(
            "Delete an account, specified by UUID, Email, Username"
        ),
    )
    delete_parser.add_argument("identifier", type=str)

    set_parser = subparsers.add_parser("set", help=("Change account settings"))
    set_subparsers = set_parser.add_subparsers()

    set_email_parser = make_subparser(
        set_subparsers,
        "email",
        set_email_command,
        help=("Change the specified UUID account's email."),
    )
    set_email_parser.add_argument("identifier", type=str)
    set_email_parser.add_argument("email", type=str)

    set_username_parser = make_subparser(
        set_subparsers,
        "username",
        set_username_command,
        help=("Change the specified UUID user's username."),
    )
    set_username_parser.add_argument("identifier", type=str)
    set_username_parser.add_argument("username", type=str)

    set_password_parser = make_subparser(
        set_subparsers,
        "password",
        set_password_command,
        help=("Change the specified UUID user's password."),
    )
    set_password_parser.add_argument("identifier", type=str)
    set_password_parser.add_argument("password", type=str)

    search_parser = subparsers.add_parser("search", help=("Search users"))
    search_subparsers = search_parser.add_subparsers()

    search_username_parser = make_subparser(
        search_subparsers,
        "username",
        search_username_command,
        help=("Search by Username."),
    )
    search_username_parser.add_argument("username", type=str)
    
    #search_username_parser.add_argument(
    #    "account_type", type=str, choices=FIRST_PARTY_ACCOUNT_TYPES
    #)

    args = parser.parse_args()
    if len(args.__dict__) <= 1:
        # No arguments or subcommands were given.
        parser.print_help()
        exit(1)

    return args


def main():
    args = get_args()
#    client_credentials = _get_client_credentials(args.env)
#    token = _get_access_token(client_credentials, args.env, args.debug)
    args.func(client_credentials, token, args)

if __name__ == "__main__":
    main()
