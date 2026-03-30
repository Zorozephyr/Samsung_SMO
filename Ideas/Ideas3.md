Cross-Domain Energy Optimization via Spatio-Temporal Graph Neural Networks
Title: Beyond Sector Sleep: Coordinated BSS-to-RAN Energy Orchestration utilizing Graph Neural Networks in Centralized SMOs

The Research Gap: Existing O-RAN energy efficiency (EE) rApps focus locally—shutting down RU sectors based on traffic. They lack the global context of OSS technician schedules, hardware degradation, and end-to-end BSS SLAs. They treat the network as isolated nodes rather than a deeply interconnected graph.

The Innovation: A Spatio-Temporal Graph Neural Network (ST-GNN) deployed as an advanced rApp within the SMO. The graph models not just the O-RU/O-DU/O-CU nodes, but integrates data from the CIM (hardware lifespan, thermal data) and BSS (SLA penalties, scheduled maintenance). The network predicts when deep-sleep states can be synchronized with physical technician dispatch windows, optimizing both power draw and operational expenditure (OpEx) simultaneously.

Proposed Architecture: The ST-GNN rApp continuously ingests topology and telemetry from the centralized DB. When it predicts an optimal EE window, it doesn't just trigger the NFO to power down the RU; it simultaneously interfaces with the OSS Workflow Manager to schedule automated maintenance tickets (via TMF 621 Trouble Ticket) without triggering false BSS SLA breach alarms.

Recommended Reading:

IEEE Transactions on Mobile Computing: "Graph Neural Networks for Wireless Communications."

O-RAN WG10 (OAM for O-RAN) whitepapers on Energy Saving (ES) use cases.

Literature on "Spatio-Temporal GNNs for Traffic Prediction."