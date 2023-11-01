# Übung 1 Lösungen

## Auf welche Herausforderungen und Einschränkungen sind die Techniker:innen bei der Microservices-Architektur gestossen?

## Was sind die Vor- und Nachteile der beiden diskutierten Architekturen 

## Untersuche in beiden Architekturen, welche Teile des Systems ein dezentrales System darstellen und welche ein verteiltes System.

## Virtualisierung
### EC2
Amazon Elastic Compute Cloud (EC2) war einer der ersten Services[^1], die AWS angeboten hat.
EC2 stellt virtuelle Maschinen (VM) zur Verfügung, die flexibel gestartet und gestoppt werden können und nur bei Gebrauch verrechnet werden.[^2]
Durch diese Flexibilität ermöglicht EC2 die Entwicklung von Services, die auf ändernde Last reagieren (skalieren) können.
Da die VM Instanzen an zahlreichen AWS Standorten deployed werden können, können auch hohe Verfügbarkeit und oder geringe Latenzen realisiert werden, da die Applikationen geographisch redundant und nahe bei den Usern deployed werden können.[^3]
Es gibt viele verschiedene Instanztypen, z.B. für hohen Speicherbedarf oder hohen CPU-Bedarf.

[^1]: https://en.wikipedia.org/wiki/Timeline_of_Amazon_Web_Services
[^2]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html
[^3]: https://en.wikipedia.org/wiki/Amazon_Elastic_Compute_Cloud
 
### ECS
Managed container orchestration service.
