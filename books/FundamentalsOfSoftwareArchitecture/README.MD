# FSA - Fundamentals of Software Architecture 
##### An Engineering Approach

By Mark Richards and Neal Ford

## Notes:

## Self-Assessment Questions:

### Chapter 1: Introduction:
1. What are the four dimensions that define software architecture?
    - Architecture consists of the ***STRUCTURE*** combined with ***ARCHITECTURE CHARACTERISTICS***, ***ARCHITECTURE DECISIONS*** and ***DESIGN PRINCIPLES***.
2. What is the difference between an architecture decision and a design principle?
    - Architecture decision is strict rule defined how the system should be constructed. For example, we have layered architecture style, we decided to follow the closed style, any layer shouldn't be able to pass through layers without the call the next beneath. Follow this rule is not recommendation! Is rule and should be followed. On the other hand, a design principle is used to provide guidance or a recommendation from the architecture to the developer team. For example, "whenever possible, leverage async messaging between services to increase performance"
3. List the eight core expectations of a software architect.
    - Make architecture decisions, continually analyze the architecture, keep current with the latest trends, ensure compliance with decisions, diverse exposure and experience, have business domain knowledge, possess interpersonal skills and understand and navigate politics
4. What is the first Law of Software Architecture?
    - There are no right or wrong answers in architecture only trade-offs.
### Chapter 2: Architecture Thinking
1. Describe the traditional approach of architecture versus development and explain why that approach no longer works.
    - The traditional approach is very similar with the waterfall project management, where the architecture do his work about the requirements and once done he passes to the development team what should be done and don't iterate about the architecture with the development team again. This causes a unidirectional arrow of interactions, resulting in architecture decisions that will never be developed by the team and some development decisions will never change the architecture.
2. List the three levels of knowledge in the knowledge triangle and provide and examples of each.
    - Stuff you, example: is the technologies/knowledge that you use in your daily basis to do your job. Stuff you know you don't know, example: is the technologies/knowledge that you have heard about it but never used or just know that exists. Stuff you don't know you don't know, example: is all the technologies/knowledge that could solve the problem that you are facing, but you even know that such thing exists.
3. Why is it more important for an architect to focus on technical breadth rather than technical depth?
    - Because an architect must make decisions that match capabilities to technical constraints. For example is much more important to an architect to know that exist 7 types of cache products and their uses cases and pros and cons, that be a specialist at one cache product, each project has different requirements and problems to solve, so he must know each product fits better in these constraints.
4. What are some of the ways of maintaining your technical depth and remaining hands-on as an architect?
    - Proof-of-concepts(POCs), tackling the low priority user stories or work on bug fixes both in an iteration and do code reviews.
### Chapter 3: Modularity
1. What is meant by the term connascence?
    - "Two components are connascent if a change in one would require the other to be modified in order to maintain the overall correctness of the system" - Meilier Page-Jones(1996-Dorset-House)
2. What is the difference between static and dynamic connascence?
    - Static connascence refers to the source-code level coupling, for example multiple components depends on a specific entity type. And dynamic connascence refers to the connascence in the execution time, for example the order execution of a method is a dynamic connascence.
3. What does connascence of type mean? Is it static or dynamic connascence?
    - Is a static connascence. Connascence of type is when multiple components depends on a specific entity type in order to work
4. What is the strongest form of connascence?
    - Dynamic connascence of type identity, when multiple system must reference to the same entity in the execution time.
5. What is the weakest form of connascence?
    - Static connascence of type name, when multiple components must agree on the name of an entity.  
6. Which is preferred within a code base-static or dynamic connascence?
    - Static connascence, because the static connascence is much more in the control of the architect and developer team to refactor to a better type of coupling than dynamic, nowadays, the lack of tools to measure dynamic connascence is high and analyze e measure dynamic connascence is much harder than static.
### Chapter 4: Architecture Characteristics Defined
1. What three criteria must an attribute meet to be considered an architecture characteristic?
    - Specifies a nondomain design consideration; Influences some structural aspect of the design; Is critical or important to application success.
2. What is the difference between an implicit characteristic and an explicit one? Provide an example of each
    - Implicit architecture characteristics is the ones that rarely appear in the requirements document, yet they are necessary for project success. For examples, security no one wants a vulnerable application but no one specifies how secure should be the app.
3. Provide an example of an operational characteristic.
    - Operational characteristic is all characteristics that a system will have to ensure responsiveness to the client an example: Recoverability
4. Provide an example of a structural characteristic.
    - Structural characteristic is all characteristics that ensure code quality an example: Maintainability and Extensibility
5. Provide an example of a cross-cutting characteristic.
    - Cross-cutting characteristic is all characteristics that ensure security an example: Privacy and Authorization.
6. Which architecture characteristic is more important to strive for availability or performance?
    - It depends on the requirements, if for the business is more important the system is need to be available 24/7 in case in any failure, so we should strive for availability, but if it is capacity of low response time in any case should be performance.
### Chapter 5: Identifying Architecture Characteristics
1. Give a reason why it is good practice to limit the number of characteristics("-ilities") an architecture should support.
    - If an architecture does all characteristics she does nothing very well, this will lead in a generic solution means that architecture will not success in the core requirements for the business. This repents because each architecture characteristics impact on each other for example when you choose to support security probably you will have to trade of by performance because you probably you're going to have to do more encryption for each data response because encryption is expensive and takes time.
2. True or False: most architecture characteristics come from business requirements and user stories.
    - False, most architecture characteristics comes from the architect interpretation of the requirements document or business stakeholder conversation. 
3. If a business stakeholder state that time-to-market(i.e., getting new feature and bug fixes pushed out to users as fast as possible) is the most important business concern which architecture characteristics would the architecture need to support?
    - Testability+deployability+agility.
4. What is the difference between scalability and elasticity?
    - Scalability is the ability to handle a large number of concurrent users without serious performance degradation, therefore elasticity is the ability to handle bursts of requests. Elasticity look for metrics like burst of users on time and scalability looks for performance of concurrent users. Some time a system can be Scalable but not elastic to handle bursts requests.
5. You find out that your company is about to undergo several major acquisitions to significantly increase its customer base. Which architectural characteristics should you be worried about?
    - Scalability the ability to handle many concurrent users on time. Also, I would be worried with Availability and deployability as architect I want to ensure that my system will not stop delivery feature because the amount of new users and the development process should not change because of that.
### Chapter 6: Measuring and Governing Architecture Characteristics
1. Why is cyclomatic complexity such an important metric to analyze for architecture?
    -  Cyclomatic complexity is metric to analyze code structure complexity where is verified cyclomatic complexity of a code using cc = E-N+2P where E represents edges(if statements count) and N represents nodes (lines of code) and P represents the number of reference for that method or function. Using this metric the architects and developers can common agree if a class or function is too complexity or not, too much complexity can mean that some architecture characteristics is not been satisfied such as testability or modularity.
2. What is an architecture fitness function? How can they be used to analyze an architecture?
    - Any mechanism that provides an objective integrity assessment of some architecture characteristic or combination of architecture characteristics. They can be used do ensure compliance of an architecture characteristic for example we could write a unit tests verifying the in dependency between architecture layers and if they are satisfying the architecture design expected. Others examples, metrics, chaos engineering and monitors.
3. Provide an example of an architecture fitness function to measure the scalability of an architecture.
    - You should define performance of your application and a range of it, for examples performance for me is the response time of each endpoint exposure in the API and I will get the simple median from this and say it is the overall performance of my application for examples high performance 300 ms and low performance 1000 ms. Ok now I will join a metric of concurrent users, such as quantity of request on small range of time and will define a monitor to verify my median performance by the concurrent users metric, if this metric reaches some kind of value of low performance if high concurrent users I will make this monitor notify my team to take a look in the system.
4. What is the most important criteria for an architecture characteristic to allow architects and developers to create fitness functions?
    - Choose the architecture characteristics that your application will satisfy the most and by fitness functions ensure the compliance.
### Chapter 7: Scope of Architecture Characteristics
1. What is an architecture quantum, and why is it important to architecture?
    -  Architecture quantum is - "An independently deployable artifact with high functional cohesion and synchronous connascence", where "independently deployable artifact" means all structure necessary to say that a system is health for examples (a monolith and his DB), and "high functional cohesion" means how well contained code is unified in purpose, and "cohesion and synchronous connascence" means sync calls within an application context or between distributed services(synchronous connascence = coupling between services in execution time by requests).
    - Architecture quantum is important because give the architects a new metric to analyze the high level architecture design, many architecture quantum indicates a distributed system, and we can assume a lot just know this kind of number. For example if a system only has 1 architecture quantum certainly this system is a monolith application
2. Assume a system consisting of single user interface with four independently deployed services, each containing its own separate database. Would this system have a single quantum or four quanta? Why?
    - Four quanta, because each separated deployable unit artifact, the independently deployed services, satisfies one quantum each. Someone also could say that this architecture is five quantum, counting the frontend as one quantum, but the front end is not an independently deployed system he depends on four services to work on, so he doesn't count in the architecture quantum count.
3. Assume a system with an administration portion managing static reference data (such as product catalog and warehouse information) and customer facing portion managing the placement of orders. How many quanta should this system be and why? If you envision multiple quanta, could the admin quantum and customer facing quantum share a database? If so, in which quantum would the database need to reside?
    - If we assume the admin system and the customer system are independently deployable units therefore this architecture has two quanta. Each has distinct responsibilities and could evolve independently. While the two quanta could technically share a database, it would be better for each to have its own to maintain the principle of independent deployability. Sharing a database would introduce a level of coupling that could inhibit the ability to deploy and scale each quantum independently. If they were to share a database, it would reside in neither quantum specifically, but would become a shared resource, which would need to be carefully managed to avoid conflicts and ensure consistency.
### Chapter 8: Component-Based Thinking
1. We define the term component as a building block of an application something the application does. A component usually consist of a group of classes or source files. How are components typically manifested within an application or service?
    - Components are usually manifested within an application or service as distinct units of functionality. They can be organized as modules, packages, or services, each encapsulating a set of related functionalities. In object-oriented languages, a component can comprise multiple classes, interfaces, or packages. In a microservices architecture, each microservice can be considered a component.
2. What is the difference between technical portioning and domain partitioning? Provide an example of each.
    - Technical portioning is a way to organize the source code by technical capabilities, for example a layered architecture presentation layer, business layer, persistence layer and so on. It makes clear that all code to lead with the database is inserted in the persistence layer. The domain portioning is a way to organize the source code by workflows, or business domain functionalities, where each part is bounded context of the system domain, for examples microservices is a architecture style where each service has his own domain. The main difference between this portioning is the top level organization.
3. What is the advantage of domain partitioning?
    - Partitioning a system by domains better reflects the kinds of changes that system will need when business rules change, models the system more close how the business functions work and more easy to migrate to a distributed architecture.
4. Under what circumstances would technical partitioning be a better choice over domain partitioning?
    - Technical partitioning is often a patter used to starter projects or business that don't have all knowledge of his domain, offering a low learning curve for developer to iterate in the code base and learn about the business. Also in an organization where exists specialists teams like the front end team, the backend team and DB teams this kind of organization better fits into technical partitioning.
5. What is the entity trap? Why is it not a good approach for component identification?
    - Entity trap is antipattern when an architect incorrectly identifies the database relationships as workflows in the application, this correspondence that rarely manifests in the real world. Components created by this pattern will result in coarse-grained components where it will not give any useful guidance to developers team.
6. When might you choose the workflow approach over the Actor/Actions approach when identifying core components?
    - By the book you should choose workflow approach when you are not in the team that not use waterfall approach to software development but in my opinion this not justify enough to say that Actor/Actions has something similar to waterfall. I would say that you should choose workflow approach when you have a good knowledge about the domain or a business specialist very close to you and when you don't have this knowledge you should go Actor/actions approach, its more easy to start from nothing identifying actor and actions.
### Chapter 9: Architecture style
1. List the eight fallacies of distributed computing
    - The network is reliable, latency is zero, bandwidth is infinite, the network is secure, topology never changes, there is only one administrator, transport cost is zero, the network is homogeneous.
2. Name three challenges that distributed architecture have that monolith architecture don't.
    - By the book: Distributed logging, distributed transactions and contract maintenance and versioning. You can have a distributed system without distributed transactions but is very hard to do that you should avoid distributed transactions more than nothing, you also have contract maintenance and versions as challenge in monolith too.
3. What is stamp coupling?
    - Stamp coupling is data structure coupling, in a distributed architecture when a system needs to call other system to fill a request data, but the requests only needs one or two properties of the response and in these properties we can say the requested has a stamp coupling on those properties.
4. What are some ways of addressing stamp coupling?
    - Create private and specialist endpoints to only return the needed data, or use other high level protocol such as GraphQL. Regardless of the technique used, we should ensure the minimal amount of data is passed between services or distributed systems.
### Chapter 10: Layered Architecture Style
1. What is the difference between an open layer and a closed layer?
    - Each layer in the layered architecture style can be either closed or open. Closed layers mean that as a request moves top-down from layer to layer, the request cannot skip any layer, but rather must go thought the layer immediately below it to get to the next layer. And in the open architecture style this restriction is not needed, and the top layer levels can jump some layer to get what is needed, for example: the presentation layer only needs DB data and nothing more, in the open style probably the presentation layer would call the persistence layer directly, jumping the business layer.
2. Describe the layers of isolation concept and what the benefits are of this concept.
    - The layer of isolation concept means that changes made in one layer of the architecture generally don't impact or affect components on other layers, providing the contracts between those layers remain unchanged. Each layer should be Independent of the others having no knowledge of the implementation of the other layer.
3. What is the architecture sinkhole antipattern?
    - Is when you have a request that pass through all layers presentation, business and persistence and database layers and only thing that happened is the data retrieve by the persistence layer, so in this call the business layer is unnecessary resulting in both memory consumption and performance.
4. What are some main architecture characteristics that would drive you to use a layered architecture?
    - Company context: As starting point when it is not known yet exactly which architecture style will ultimately be used. Also is a good architecture for being startups that don't have so much knowledge about the business that they are building.
    - Architecture characteristics: Development cost and learning curve resulting in a very simple architecture. 
5. Why isn't testability well support in the layered architecture style?
    - Because the author simple said that as a monolith application get a lot of changes this makes test difficult. I totally disagree on this one, a monolith application with layered architecture style only became difficult to test when you poor isolation levels between layers and a high level of coupling between entities. But, the case of one front end, one backend and one database are the best scenarios to make integrations tests and end-to-end tests, as you start adding third party APIs into this monolith this kind of testing became very difficult to do.
6. Why isn't agility well-supported in the layered architecture style?
    - Because deployability is not a very well-supported characteristic, for every change that needs to be release you have to change all available instances, and as the authors assumed that testability is not well-supported to the sum of hard to deploy and hard to test will result in an architecture that slow down the teams productive.
### Chapter 11: Pipeline Architecture
1. Can pipes be bidirectional in a pipeline architecture?
    - It is possible to make bidirectional pipelines, but the recommendation is to make unidirectional and point-top-point for performance, simplicity that will encourage composition.
2. Name the four types o f filters and their purpose.
    Filters are self-contained, independent of other filters and generally stateless, composite tasks should be handled by a sequence of filter rather than a single one.
    - Producer: The starting point of a process, outbound only, sometime called the source. Transformer: Accepts input, optionally performs a transformation on some or all of the data, then forwards it to the outbound pipe. Tester: Accepts input test one or more criteria, then optionally produces output, based on the test. Consumer: The termination point for the pipeline flow. Consumer sometimes persist the final results of the pipeline process to a database, or they may display the final results on a user interface screen.
3. Can a filter send data out through multiple pipes?
    - Yes, in the book example workflow show us the `DurationFilter(tester)` Sending data to `DurationCalculator(transformer)` and to `UptimeFilter(tester)` what indicates multiples pipes.
4. Is the pipeline architecture style technically partitioned or domain partitioned?
    - Technical partitioned
5. In what way does the pipeline architecture support modularity?
    - Ensure each filter has a purpose, one directional data flow, self-contianed and independent of others are all characteristics encourages modularity and separation of concerns between the various filters types and transformers.
6. Provide two examples of the pipeline architecture style.
    - ETL(extract, transform and load) - a common patterns to data lakes with many sources of data. Apache camel.
### Chapter 12: Microkernel Architecture
1. What is another name for the microkernel architecture style?
    - Plug-in architecture. 
2. Under what situations is it ok for plug-in components to be dependent on other plug-in components?
    - There is no situation where this is ok, when you are doing this you are breaking the core idea behind this architecture style design a solid core system where plug-ins are literally just plug-ins and can have their own domain. 
3. What are some of the tools and frameworks that can be used to manage plug-ins?
    - Open service gateway(Java), Penrose(Java) and Prism(.Net).
4. What would you do if you had a third-party plug-in that didn't conform to the standard plug-in contract in the core system?
    - Probably if is a third-party plug-in is others team, or companies trying to use the core system develop by my team. So assuming this conditions, probably I would develop a system where have simple public interface and very well written messages errors for integration for this interface, providing system that any developer can fix what is wrong with the integration with our core system. If the public interface is not satisfied the integration cannot happen.   
5. Provide two example of the microkernel architecture style.
    - Monolith microkernel and distributed microkernel where plug-ins can be segregated deployable units. Examples of usages: Eclipse IDE and Jenkins
6. Is the microkernel architecture style technically partitioned or domain partitioned?
    - Microkernel can be both, technically partitioned or domain partitioned.
7. Why is the microkernel architecture always a single architecture quantum?
    - (By the book →) Because all requests must go through the core system to get to independent plug-in components. (Explaining why the authors are saying this) As the definitions of quantum says "An independently deployable artifact with high functional cohesion and synchronous connascence", if you have microkernel architecture style where each plug-in can be an "An independently deployable artifact" and also has his own database and as architecture suggest the plug-in must be independent plug-in components satisfying the only thing that this style don't satisfy very well is the "high functional cohesion" because has to pass through the core system first to reach de plug-ins and just because of that the authors says that always will be a single quantum. But in my opinion if you choose use a microkernel architecture with distributed plug-ins components you're choosing almost all trade-offs of microservices architecture style, consider this style as a transitioning one to the another styles. 
8. What is domain/architecture isomorphism?
    - Architecture isomorphism = "how close toes the shape of the architecture match the shape of the problem domain?". How you translate very well the problem domains to the right architecture.
### Chapter 13: Service-Based Architecture
1. How many services are there in a typical service-based architecture?
    - Because the services typically share a single monolith database, the number of services within an application context generally range between 4 and 12.
2. Do you have to break apart a database in service-based architecture?
    - Is not required, usually they share a single monolith database. This database coupling represents an issue where it is not done properly can a table schema change can result in an impact to every service. 
3. Under what circumstances might you want to break apart a database?
    - You want to change to a microservice architecture and your services are very well-defined so the entities tables don't have many dependencies making easy to migrate to a separate database server.
4. What technique can you use to manage database changes within a service-based architecture?
    - Using a single library for each entity schema making the impact of a change only in the services that are using the schema changed.
5. Do domain services require a container (such as docker) to run?
    - No.
6. Which architecture characteristics are well-supported by the service-based architecture style?
    - Modularity, overall cost, fault tolerance, testability, deployability and reliability. 
7. Why isn't elasticity well-supported in a service-based architecture?
    - As the number of services increase the database server start to be the bottleneck to handle too many concurrent users request, more functionality is replicated than with finer-grained services (such as microservices) and as such is not as efficient in terms of machine resources and not as cost-effective.
8. How can you increase the number of the architecture quanta in a service-based architecture?
    - Increasing the number of user interfaces by domain, each of domain will represent an amount of backend services and one database resulting in many quanta.
### Chapter 14: Event-Driven Architecture Style
1. What are the primary difference between the broker and mediator topologies?
    - Centralization of event management, the broker differs from the mediator in that there is no central event mediator.
2. For better workflow control would you use the mediator or broker topology?
    - The mediator topology, this topology can handle a required order of event broadcast and event consuming by ensuring that such order of events broadcasted are published with success and the mediator can move on in the orchestrations and coordination of events.
3. Does the broker topology usually leverage a publish-and-subscribe model with topics or a point-to-point model with queues?
    - Publish-and-subscribe model with topics, in this topology each service broadcast his own events and the services in interests in such topic should process the event.
4. Name two primary advantage of asynchronous communications
    - Increase the overall responsiveness and decoupling.
5. Given an example of typical request within the request-based model.
    - A HHTP request using REST pattern, for example, the order-service must include in his response all products names and cost by each product in a request Get by ID order, so the order-service request to the product-service by REST request. (I don't know if I should include path and response examples here.)
6. Given an example of typical request in an event-based model.
    - A service input a message in a queue or an event in a topic and in this topic or queue has one purpose explicit a necessity, we're going to have two queue or topics one to request others to response, so each service can ask each other data by the queue or topics.
7. What is the difference between an initiating event and a processing event in event-driven architecture?
    - Processing event occurs when other service has already emitted an event, so the consumer of an event will start a process by this trigger and in initiating usually means that a service finish a process and would like to notify his alterations or ask other services to start the process by event. 
8. What are some of the techniques for preventing data loss when sending and receiving messages from a queue?
    - For a service he should persist the messages not juts in memory but also in a physical data store(fileSystem, database, others...) and to brokers when receive a message he should always persist this message in a physical data store too so when he goes down he can ask what message were in the queue at the moment that he crashed. A broker should just dequeues a message only if the consumer service response with success the message received.
9. What are three main driving architecture characteristics for using event driven architecture?
    - Fault tolerance, Performance and Scalability
10. What are some of the architecture characteristics that are note well-supported in event driven architecture?
    - Simplicity and testability. Maintaining an entire system built with event communication is harder than a request based model due the necessity to integrate to many systems, nondeterministic and dynamic event flows, also test the integration between services is harder, you only will achieve some kind of automation test to test integration between services in test environment deployed, to have this kind of tooling is expensive and hard to maintain.
### Chapter 15: Space-Based Architecture
1. Where does space-based architecture get its name from?
    - Tuple space, a technique of using multiple parallel processors communicating through shared memory.
2. What is primary aspect of space-based architecture that differentiates it from other architecture styles?
    - Couple of components dedicated only to the architecture works, we have a process unit, in memory data grid, data replication engine, data reader, data writer, deployment manager, data grid, processing grid and messaging grid all these components are necessary to support the idea of shared memory and prevent to access database directly providing an enormous performance, scalability and elasticity.
3. Name the four components that make up the virtualized middleware within a space-based architecture.
    - Deployment manager, data grid, processing grid and messaging grid
4. What is the role of the messaging grid?
    - Manages the input request and session state. When a request comes into the virtualized middleware the messaging grid determines which active processing components are available to receive the request and forwards the request to one of those processing units.
5. What is the role of a data writer in space-based architecture?
    - Receive messages (async communication) from the data pumps and process units and updates the database state.
6. Under what conditions would a service need to access data through the data reader?
    - A crash of all processing units instances of the same named cache, a redeployment of all processing units withing the same named cache, or retrieving archive data not contained in the replicated cache
7. Does a small cache size increase or the decrease the chances of a data collision?
    - Decrease due the inversely proportional relationship between cache size and Collision Rate in the equation CR = N * ((U*Rˆ2)/S)*RL where S is the size of the cache, number of rows.
8. What is the difference between a replicated cache and a distributed cache? Which one is typically used in space-based architecture?
    - Replicated cache is when each processing unit(instances) has in memory data cache and distributed cache is when you delegate to a dedicated server cache to do this task, so your processing units became stateless. The decision about each strategy to use will depends on the requirements of the system if you're prioritizing consistency is better goes to with distributed, also with you leading with a large amount of data to cache certainly you will need to use the distributed strategy due the cost of use high memory instances. So my recommendation would be, build your processing units agnostic of the type cache they are using if is replicated or distributed shouldn't matter to the process, starts with replicated when all things a small, and you have low Collision Rates, when the Collision Rates starts to grow consider to change to distributed cache server.
9. List three of the most strongly supported architecture characteristics in space-based architecture.
    - Performance, Scalability and Elasticity.
10. Why does testability rate is so low for space-based architecture?
    - In terms of unit testing is not a problem, but perform integration testes and have environment ready to test all conditions its complex due the cache, not cached data and also the data async writing, other aspect to test is hard is test the high levels of scalability and elasticity supported in this architecture. Support tooling to make all this process fast will be expensive to develop and also to maintain.