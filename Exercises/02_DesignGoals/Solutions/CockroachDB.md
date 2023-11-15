## Fallstudie Design Goals: Cockroach DB
### Why use CockroachDB?
There are many reasons to use CockroachDB, including:
- Resiliency
- Scalability
- Strong consistency
- Geo-partioning and multi-region features
- PostgreSQL-compatibility

### Diskussion
#### Resource sharing
Ein Datenbanksystem kann als solches als gemeinsam genutzte Resource angesehen werden.
Da es meistens nicht sinnvoll ist, für jede Applikation erneut eine Persistenzschicht zu implementieren werden Daten, die persistiert werden müssen in einer Datenbank gespeichert.
Wenn sich diese Datenbank auf einem entfernten System befindet, handelt es sich um ein verteiltes System.

CockroachDB selber nutzt auch eine verteilte Architektur um verteilte Compute und Storage Ressourcen von Cloud Anbietern gewinnbringend zu nutzen.

#### Distribution Transparency
CockroachDB unternimmt einiges, um die verteilte Architektur zu verstecken.
Requests werden zum geographisch naheliegendsten Node geleitet, sowohl schreibend als auch lesend.
So bleiben die Latenzen tief, auch wenn die Nodes global verteilt sind.

#### Openness
CockroachDB ist zum PostgreSQL Dialekt kompatibel und ist so sehr offen.
Es können bestehende PostgreSQL Datenbanktreiber verwendet werden.
Und, im Rahmen des implementierten Sprachumfangs, kann die Datenbank mit einer anderen PostgreSQL kompatiblen Datenbank ausgetauscht werden.

#### Dependability
Zuverlässigkeit ist ein zentrales Feature von CockroachDB.
Die Daten werden mehrfach repliziert, bei einem Ausfall von einem Node sind die Daten weiterhin verfügbar.
Ist ein Node wieder online, werden die Daten automatisch wieder repliziert.
![](maintain-availability.png)

#### Security
CockroachDB unterstützt gängige security features.

#### Scaling
CockroachDB kann mit einer theoretisch beliebigen Anzahl an Nodes betrieben werden, auf welche die Last verteilt wird.
So kann in sehr grossem Ausmass skaliert werden.
Möglich ist dies, weil die einzelnen Nodes, anders als bei traditionelleren Datenbanksystemen, sowohl schreibend als auch lesend arbeiten können.