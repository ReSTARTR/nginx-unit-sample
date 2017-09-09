def application(environ, start_response):
    start_response('200 OK', [('Content-type', 'text/plain')])
    return 'Hello, Python'

from wsgiref import simple_server

if __name__ == '__main__':
    server = simple_server.make_server('', 8400, application)
    server.serve_forever()