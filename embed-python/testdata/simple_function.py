def get_params(param_url, ab_version):
    """This is the simple function test.

    get_params takes param_url as a dict input, ab_version as an integer input,
    and it will output the JSON string for server to parse and respond to the client.

    >>> print(get_params({"cache": 1, "ab_version": "10.0.0.194"}, 203))
    {"cache": 1, "ab_version": 203}
    >>> print(get_params({"ab_version": "10.0.0.194"}, 203))
    {"cache": 10, "ab_version": 203}
    >>> get_params({"ab_version": "10.0.0.194"}, 203)
    '{"cache": 10, "ab_version": 203}'
    """
    cache = param_url.get("cache", 10)
    version = ab_version
    return '{"cache": %d, "ab_version": %d}' % (cache, version)

# python3 -m doctest simple_function.py
print(get_params({"ab_version": "10.0.0.194"}, 203))

