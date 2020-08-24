# Dummy Go App using GQLGEN


Schema.graphqls - basically where i define the structure or what is to be returned from API. This is the first thing I do.

Model_gen - This gets generated from the above schema.

Schema_resolver - for all your functions

Postgres - I create Db and call it

pqueries - is where i define db functions

getLogic - uses the pqueries func
