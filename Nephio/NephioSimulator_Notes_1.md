Regional Represents the cloud cluster and control plane
Edge01 represents User plane and closer to end user cluster

UE(End User) -> End User.In a lab/sandbox, this is usually a software simulator (like UERANSIM) pretending to be a 5G smartphone.

gNodeB(5G base station) -> 5G Base Station. In a lab/sandbox, this is usually a software simulator (like Free5GC) pretending to be a 5G base station.

UPF(user plane function) -> high speed router, SIM -> GNB -> UPF -> Internet. UPF tries to make this flow as fast as possible

AMF(Access and Mobility Management Function) -> Authentication, Authorization, Registration, Mobility...So when you turn ur phone on, it will decide if ur phone is allowed to connect to the network

SMF(Session Management Function) -> It sets up tunnel for the user. This user is watching a video; please route their traffic to the internet

AUSF (Authentication Server Function): Security. It runs the complex cryptography to prove the SIM card is valid.

NRF (Network Repository Function): The Phonebook. It keeps a list of all running network functions so they can find each other (e.g., the AMF asks NRF, "Where is the SMF?").

NSSF (Network Slice Selection Function): The Slicer. If the network is cut into "slices" (one for gaming, one for IoT), this component decides which slice the user gets put into.

UDM (Unified Data Management): The Frontend Database. It processes user data (subscription info, allowed services).

UDR (Unified Data Repository): The Backend Storage. The actual database where the UDM and PCF store their data.

PCF (Policy Control Function): The Lawyer. It decides the rules: "User X has a Gold plan, give them high speed," or "User Y is out of data, throttle them."

WebUI: This is likely a specific dashboard for the free5gc lab to visualize subscribers; it is not a standard 3GPP network function.