# Python module containing example query routing table function.
# Configure path/name to this file in [pgbouncer] section of ini file. 
# Ex:
#    routing_rules_py_module_file = /etc/pgbouncer/routing_rules.py


# ROUTING TABLE
# ensure all dbkey values are defined in [database] section of the pgbouncer ini file 
# Test by calling routing_rules() with sample queries, and validating dbkey values returned
routingtable = {
	'route' : [{
			'usernameRegex' : '.*',
			'queryRegex' : '(?i)SELECT.*',
			'dbkey' : 'example.2'
		}, {
			'usernameRegex' : '.*',
			'queryRegex' : '(?i)(INSERT|UPDATE|DELETE|START|COMMIT|ROLLBACK).*',
			'dbkey' : 'example.1'
		}
	],
	'default' : None
}


# ROUTING FN - CALLED FROM PGBOUNCER-RR - DO NOT CHANGE NAME
# IMPLEMENTS REGEX RULES DEFINED IN ROUTINGTABLE OBJECT
# RETURNS FIRST MATCH FOUND
import re
def routing_rules(username, query):
	for route in routingtable['route']:
		u = re.compile(route['usernameRegex'])
		q = re.compile(route['queryRegex'])
		if u.search(username) and q.search(query):
			return route['dbkey']
	return routingtable['default']

if __name__ == "__main__":
    print "test for insert: " + routing_rules("postgres", "insert values ('name') into tablea;")
    print "test for update: " + routing_rules("postgres", "update tablea set name='bla' where x='y';")
    print "test for delete: " + routing_rules("postgres", "delete from tablea;")
    print "test for start transaction: " + routing_rules("postgres", "start transaction;")
    print "test for select: " + routing_rules("postgres", "select * from tableb;")
    

