# Note
This project is only to illustrate the general ideas of the processes I come up with. It's been aware that real-life scenario is much more intricate. Integrating 3rd-party api would be out-of-scope with this mini demo. The code itself has no real-life value and is not intended to be used in any prod env. 

## Process overview
The user shall first drop a query. The server first check if an existing query is similar (using faiss db e.g.) to the coming one and whether the location between two users are close enough to share the results. Also, the previous query should not be overdue.

If all checked then retrun the old result. Otherwise, run the api and get new results, store them in the db and return the new response to the user.

Once the new results is returned, store the results and embed the query and save them for future reference. 

## Further plan
- High availability. In order for users to have better access to the system, structures including microservices and load balancer should be considered. (go-zero, nginx, HAproxy,docker, etc. )
- Distributed transaction. In order to ensure the consistency of the system, distributed transactions should be considered. (MQ, etc. )
- Expandability. A simple change in query might result in a big difference in the way Transformer interacts with an API. More agents(ways of interacting with APIs) should be considered. (e.g. a new agent for a new API)

