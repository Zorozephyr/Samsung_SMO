Idea 4: Multi-Agent Synthesis for NFO Workflows
Title: Agentic Orchestration: Dynamic DAG Synthesis in O-RAN SMO using Multi-Agent Frameworks

The Research Gap: Centralized NFOs rely on static, pre-programmed workflows (DAGs) to handle network configurations or fault remediations. When a novel, multi-cluster failure occurs, the static rules fail, requiring manual intervention.

The Innovation: Replace rigid workflow templates with a multi-agent AI system. Different agents assume specialized roles (e.g., a "Telemetry Agent" parsing the Centralized DB, a "Standards Agent" holding 3GPP rules, and a "Synthesis Agent"). These agents collaborate in real-time to generate a custom, executable YAML/JSON workflow on the fly to hand off to the NFO.

Proposed Architecture: When the Workflow Manager detects an anomaly, it triggers the multi-agent orchestration layer. The agents query the CIM for live cluster state, debate the optimal recovery path, and construct a deterministic DAG. The NFO then executes this generated workflow. This brings resilient, persistent-memory AI workflows directly into the telecom control plane.

Recommended Reading:

Research on LLM Multi-Agent Frameworks for IT Operations (AIOps).

O-RAN WG2 (Non-RT RIC & A1 Interface) specifications on AI/ML lifecycle management.