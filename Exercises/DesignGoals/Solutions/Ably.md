# Lösungen
## Welches sind die wichtigsten Features?
- Höchste Datenintegrität
  - Mindestens 8x9 (99.999999%) message survivability on failures
  - Exactly-once delivery
  - 100% guaranteed message delivery
- Availability
- Höchste Verfügbarkeit: 99.999% uptime SLA
- Enorm grosser Datendurchsatz möglich
- Sehr niedrige Latenzen von <100ms global

[^1]: https://ably.com/four-pillars-of-dependability

## Für welche von diesen Features werden verteilte Systeme eingesetzt?
Der Service ist auf 7 Rechenzentren und 385 Points of presence (PoP) verteilt.
Dadurch sind global immer nahe Standorte verfügbar, die die Nachrichten mit kurzer Latenz weiterleiten können.
Dies ermöglicht auch die sehr hohen Verfügbarkeitsgarantien und hohe Fehlertoleranz bei Ausfällen in den Rechenzentren.
Weiter können dadurch auch ein sehr hoher Datendurchsatz erreicht werden, da das System in grossem Masse Rechenzentrum übergreifend horizontal skaliert werden kann.

Datendurchsatz wird per channel limitiert, wodurch global unlimitierter Durchsatz erreicht werden kann.

## Wie wird dies umgesetzt?
Im Detail ist dies nicht klar, da es ein proprietäres System ist.

## Welche der folgenden Eigenschaften sind **Transparent**, welche nicht? 
- Zugriff: Vollständig Transparent, die zugrunde liegenden Betriebssysteme und deren Zugriffsarten sind abstrahiert
- Standort: Vollständig Transparent, der Standort des Clients oder des Servers ist nicht relevant und nicht ersichtlich.
- Replikation: Vollständig Transparent, die Replikation der Daten geschieht vollständig automatisch und transparent
- Änderung des Standorts: 
- gleichzeitiger Zugriff von verschiedenen Usern: Vollständig Transparent, die Loads der anderen User haben keinen Einfluss auf die eigene Nutzung
- Fehler: Vollständig Transparent, Ausfälle in Rechenzentren führen in den seltensten Fällen zu einer Beeinträchtigung des Service

## Ist das System **offen**?
Teilweise. Als propriertäres System ist es nicht auf Code-Ebene erweiterbar.
Jedoch bietet das System zahlreiche Integrationen zu anderen Systemen.[^2]

[^2]: https://ably.com/integrations

## Wie **zuverlässig** ist das System?
Sehr zuverlässig. Zuverlässigkeit ist ein Kern-Feature.

## Wie wird die **Sicherheit** gewährleistet?
Nachrichten können mit eigenen Schlüsseln end-to-end verschlüsselt werden.
Dadurch ist die Nachricht während sie im Ably System ist dauernd verschlüsselt mit einem Schlüssel, den Ably nicht besitzt.[^4]

Es werden verschiedene Authentifizierungmethoden unterstützt.[^5]

[^4]: https://ably.com/docs/channels/options/encryption?lang=javascript
[^5]: https://ably.com/docs/auth

Für Enterprise Customers kann der Datenfluss geographisch eingeschränkt werden.

## In welchem Rahmen kann das System **skalieren**? Wie wird dies umgesetzt?
Skalierbarkeit ist ein zentrales Feature.
Die konkrete Implementierung ist nicht ersichtlich.

Es werden Limitierungen auf einzelnen Channels gemacht, wodurch das System global vorhersagbar wird.

Durch die Verteilung auf mehrere Rechenzentren und PoPs können Lastschwankungen global verteilt werden.
Durch die Verwendung von AWS Cloud Infrastruktur kann ein grosser Teil der Skalierung auf den Cloud Provider abgewälzt werden. [^3]

[^3]: https://ably.com/aws