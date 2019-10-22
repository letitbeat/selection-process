# Selection Process

Selection Process is composed by two microservices `api` and `scoring`

- `api` : Exposes an endpoint `/task/{taskId}` which returns a list of applicants scored by tags for a given task.
- `scoring` : Exposes an endpoint `/scoring/{taskId}` which calculates the scores for a requested list of applicants for the given task.

## How to use it

First, we need to build and run the MongoDB and the services defined in the `docker-compose.yml` file.

To achieve this, we run the following command inside the root folder:

```
make run
```

The output should be something similar to this:
```
$ make run
docker-compose up -d
Creating network "selection-process_default" with the default driver
Creating selection-process_api_1           ... done
Creating selection-process_scoring_1       ... done
Creating mongo                             ... done
Creating selection-process_mongo-express_1 ... done
```

After this step, we proceed to import our test data files, which should be located at `/resources/dump`

To do this we run:
```
make import
```
The output of the previous command should be something similar to this:
```
/bin/bash ./scripts/importer.sh
2019-08-12T22:01:21.130+0000    preparing collections to restore from
...
...
2019-08-12T22:01:21.205+0000    finished restoring main.tasks (100 documents)
2019-08-12T22:01:21.214+0000    no indexes to restore
2019-08-12T22:01:21.215+0000    finished restoring main.users (1098 documents)
2019-08-12T22:01:21.215+0000    done
```

After successfuly imported our data we can proceed to test the `task` endpoint by open this URL in the browser:

http://localhost:8080/tasks/004DFGCdNuyccjCHQ

And should be able to see something similar to this:

```javascript
{
    "taskId": "004DFGCdNuyccjCHQ",
    "applicants": [
        {
            "siderId": "7Ba7AetuSxRYuDWoj",
            "firstName": "Andreas",
            "lastName": "Macha",
            "score": 100
        },
        {
            "siderId": "9Gt4gk3iiyyrQ9NjJ",
            "firstName": "Margorie",
            "lastName": "Croom",
            "score": 100
        },    ],
    "description": "<p>Frichti a besoin de toi pour l'aider dans ses préparations de commandes.</p><h4 class=\"tpl-task-description-title\">Ton rôle</h4><p>Tri/Rangement<br />Trier, répartir et ranger les produits selon des règles pré-déterminées.</p><p>Picking<br />Prélever et rassembler les produits de manière ordonnée.</p><h4 class=\"tpl-task-description-title\">Tes objectifs</h4>",
    "country": "FR"
}
```

et voilà!

### TODO

- Implement gRPC
- Common package for model structs
- Improve error handling and logging
- async routines for scoring
