import requests

SESSION = requests.Session()

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

def main():
    params = {
        "filterKey": "username",
        "filterData": ["MyUsername", "MyUsername1"]
    }
    headers = {}
    get("http://127.0.0.1:8080/v1.0/users", params, headers, [200], "", True)

if __name__ == "__main__":
    main()