<!-- headingDivider: 4 -->
<style>
img {
  display: block;
  margin: 0 auto;
}
</style>

# Einführung

[This work is published under a Creative Commons BY-SA license.](https://creativecommons.org/licenses/by-sa/4.0/)

## «Parallele Systeme»

- Eine single core CPU kann nur einen Prozess gleichzeitig ausführen
- Multi-core CPUs entsprechend mehrere gleichzeitig
- Ausser in sehr einfachen Embedded Systemen müssen jedoch immer sehr viele Prozesse «gleichzeitig» ausgeführt werden
  können z.B. auf einem Server oder auf einem Desktop Computer
  ![img.png](Images/htop.png)

## «Parallele Systeme»

- Viele verschiedene Prozesse (tausende) werden von einem oder mehreren (bis zu dutzenden) Prozessoren ausgeführt
- Ein einzelner Prozessor kann demnach nacheinander mehrere Prozesse bearbeiten
- Die Prozessoren befinden sich auf demselben Chip oder auf dem selben Mainboard
- Sie haben geteilten sowie gemeinsamen Speicher
- Die Verbindung zwischen ihnen (Interconnect) hat geringe Latenz, hohe Bandbreite und ist zuverlässig.
  ![img.png](Images/parallelSystems)

---

- Parallele Ausführung (parallelism): Mehr als eine Aufgabe wird gleichzeitig ausgeführt
- Nebenläufig (concurrency): Mehr als eine Aufgabe wird abgearbeitet (durch schnelles context switching)
  ![img.png](Images/concurrent_vs_parallel.png)

---

- Eine zentrale Aufgabe von Betriebsystemen ist es, die Prozesse auf die CPUs zu verteilen.
- Dies wird «Scheduling» genannt.

## Verteilte Systeme

> «A distributed system is a collection of independent computers that appears to its users as a single coherent system.»

![img.png](Images/DistributedSystem.png)
VanSteen, 2017, S. 26

P: Prozessor, Interconnect: Netzwerkverbindung, meistens HTTP, UDP/TCP, IP, Ethernet basiert

# Resource Sharing

- Ressourcen verfügbar machen: Drucker, Computing, Storage, Daten, Netzwerk
- Teure Ressourcen können besser ausgelastet werden und müssen nicht mehrfach angeschafft werden
- Zusammenarbeit

## Domain Name System

![img.png](Images/DomainNameSystem)

---

![w:1000px](Images/when-the-cloud-leaves-the-datacenter-530836-1.jpg)

## Anforderungen an moderne Software

- Hohe Verfügbarkeit
- Skalierbarkeit
- Im Katastrophenfall sollen die Systeme schnell wiederhergestellt werden können
- Soll funktionieren, auch wenn Teile des Systems Offline sind (Resilienz)
- Kostengünstig
- Einfach
- Updates müssen einfach eingespielt werden können

## Lösungsansätze

- Replication: Masking Failures
- Tradeoff: Teuer und Komplex

![img.png](Images/LoadBalancing.png)

## Populäre verteilte Systeme

- [Matrix](https://element.io/features/decentralised-matrix-network)
- [Mastodon](https://joinmastodon.org/)
- [Nextcloud](https://nextcloud.com/de/)
- [CockroachDB](https://www.cockroachlabs.com/product/)
- [Neon](https://neon.tech/)
- [Ably](https://ably.com/)
- ...


# Koordination

- Tasks können gleichzeitig ausgeführt werden
- Gleichzeitiger Zugriff auf gemeinsame Daten kann in inkonsistenten Daten resultieren

## Mutex

- MUTual EXclusion: wechselseitiger Ausschluss
- Einfachste Möglichkeit, Ressourcen für alle anderen zu blockieren
- Critical Section wird mit acquire() und release() umschlossen
- acquire() und release() müssen atomare Operationen sein (Hardwareunterstützung)

### Mutex

```C
acquire() { 
  while (!available)
    /* busy wait */
  available = false;;
}

release() {
  available = true;
}
```

```C
do {
    acquire lock
       critical section
    release lock
      remainder section
 } while (true);
```

## Semaphore

- Mehr Möglichkeiten als Mutex
- Schützt gemeinsame Ressourcen
- Counting semaphore: Mehrere Ressourcen
- Binary semaphore: Nur eine Ressource
- Ein Zugriff auf eine gemeinsame Ressource wird mit dem Nehmen und Geben umschlossen

### Beispiel

![w:100%](Images/semaphore.png)

# Replication

The following slides adapted of Martin Kleppmann's Course at University of Cambridge: 
https://www.cl.cam.ac.uk/teaching/2425/ConcDisSys/materials.html

- Keeping a copy of the same data on multiple nodes: Databases, filesystems, caches, ...
- A node that has a copy of the data is called a replica 
- If some replicas are faulty, others are still accessible
- Spread load across many replicas
- Easy if the data doesn’t change: just copy it
- We will focus on data changes

## Idempotence

A function f is idempotent if f(x) = f(f(x)).

Choice of retry behaviour:
- **At-most-once semantics:** send request, don’t retry, update may not happen
- **At-least-once semantics:** retry request until acknowledged, may repeat update
- **Exactly-once semantics:** retry + idempotence or deduplication


## "Consistency"

A word that means many different things in different contexts!
- **ACID**: a transaction transforms the database from one “consistent” state to another
  - Here, “consistent” = satisfying application-specific invariants e.g. “every course with students enrolled must have at least one lecturer”
- **Read-after-write consistency**
- **Replication**: replica should be “consistent” with other replicas
  - “consistent” = in the same state? (when exactly?) 
  - “consistent” = read operations return same result?
- Consistency model: many to choose from

---

Recall atomicity in the context of ACID transactions:
- A transaction either commits or aborts
- If it commits, its updates are durable
- If it aborts, it has no visible side-effects
- ACID consistency (preserving invariants) relies on atomicity

## Strong Consistency: Linearizability

- Informally: every operation takes effect atomically sometime after it started and before it finished
- All operations behave as if executed on a single copy of the data (even if there are in fact multiple replicas)
- Consequence: every operation returns an “up-to-date” value, a.k.a. “strong consistency”

Linearizability advantages:
- Makes a distributed system behave as if it were non-distributed
- Simple for applications to use

Downsides:
- Performance cost: lots of messages and waiting for responses
- Scalability limits: leader can be a bottleneck
- Availability problems: if you can’t contact a quorum of nodes, you can’t process any operations

## Eventual Consistency

- Eventual consistency: a weaker model than linearizability. Different trade-off choices.
- Replicas process operations based only on their local state.
- If there are no more updates, eventually all replicas will be in the same state. 
- No guarantees how long it might take

## Local-first software

End-user device is a full replica; servers are just for backup. “Local-first”: a term introduced by me (Martin 
Kleppmann) and my colleagues https://www.inkandswitch.com/local-first/

Calendar app with cross-device sync as an example:
- App works offline (can both read and modify data)
- Fast: no need to wait for network round-trip
- Sync with other devices when online
- Real-time collaboration with other users
- Longevity: even if cloud service shuts down, you have a copy of your files on your own computer
- Supports end-to-end encryption for better security
- Simpler programming model than RPC
- User control and agency over their own data

# Publish-subscribe Architekturen

![w:600px](Images/distributedSystems/02/02-09.png)

---

![w:600px](Images/distributedSystems/02/02-10a.png)

---

![w:600px](Images/distributedSystems/02/02-10b.png)

---

![w:600px](Images/distributedSystems/02/02-12.png)

---

![w:800px](Images/distributedSystems/02/02-14b.png)

---

![w:1000px](Images/distributedSystems/02/02-17.png)

# Cloud und Edge Computing

> The entire history of software engineering is that of the rise in levels of abstraction.

-- Grady Booch

---

![w:1000px](Images/PizzaAsAService.png)

## Abstractions

![w:1000px](Images/Abstractions.png)
(VanSteen, 2017, S. 30)

## XaaS

![w:1000px](Images/XaaS.png)

## Fallstudie

![w:1000px](Images/CaseStudyDashboard.png)

## Edge Computing

![02-28.png](Images/distributedSystems/02/02-28.png)
