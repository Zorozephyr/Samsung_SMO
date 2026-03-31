The "Nano-NFV" (Split-Execution Data Plane)
Title: Orchestrating the Nano-NFV: BSS-Driven Split-Execution of eBPF and WebAssembly for Sub-Millisecond Ephemeral Slicing

The Research Gap: Traditional NFVs are deployed as heavy Docker containers/K8s pods. Even with optimized Helm charts, spinning up a new UPF or Deep Packet Inspection (DPI) function takes seconds to minutes. This is too slow for 6G ephemeral slicing, where a dynamic BSS order might request a slice for only a few seconds.

The Innovation: The SMO does not deploy pods. Instead, the NFO acts as a compiler. When a TMF service order arrives, the SMO synthesizes a "Nano-NFV" split into two parts: an eBPF hook (acting as the SDN router to grab packets in the kernel) and a Wasm binary (acting as the NFV to perform stateful payload inspection). Both are injected directly into the worker node's runtime in milliseconds, bypassing container orchestration entirely.

Proposed Architecture: The BSS payload hits the Workflow Manager. The NFO queries the Central Database (CIM) for target IPs. It pushes the pre-compiled Wasm/eBPF binaries to the edge node. eBPF handles the L2/L3 SDN traffic steering, instantly passing specific high-value packets into the Wasm sandbox for L7 NFV processing, tearing down instantly when the SLA expires.

Recommended Reading:

Research on Proxy-less Service Meshes (e.g., Cilium integrating eBPF and Envoy/Wasm).

Academic papers on WebAssembly for Network Function Virtualization.