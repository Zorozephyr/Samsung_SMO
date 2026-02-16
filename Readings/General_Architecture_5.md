In real world, companies dont always deploy seperate physical boxes for every function. They mix, match and merge them to save money or space.

Common Combinations:
1. Disaggregated(Seperate): Good for massive scale where you want different vendors for every part.
2. O-CU Combined: The O-CU-CP and O-CU-UP are in one box
3. The all in one small cell: The O-CU, O-DU and O-RU are all squashed into a single physical unit. This looks like traditional Wifi access point but runs on 5G.

Expert Note: Even if they are in one physical box, the O-RAN interfaces (like E2) still exist logically inside the software so the RIC can still control them


Shared Cell:
    Usually one radio unit = one cell
    But sometimes u want to create multiple radios to create a single seamless coverage area(Like a long tunnel or a stadium)

    O-RAN defines 2 ways to hookup multiple O-RU to O-DU port to create this shared cell

    Mode1: FHM(Fronthaul Multiplexer) - "The hub"
        FHM copies signal from O-DU and blasts it to all connected O-RUs at once. It combines the signals coming back up
    
    Mode2: Cascade Mode - "The daisy chain"
        O-DU -> O-RU#1 -> O-RU#2....
        Every RU in the chain has a built copy and combine to pass signal down the line

Fronthaul Gateway(FHGW):
    Sometimes u have O-RU's that dont speak the standart "Open Fronthaul 7-2x" language
    FHGW sits between DU and RU and acts as an adapter
    Interface:
        North Side (to O-DU): Speaks standard Open Fronthaul (7-2x).
        South Side (to O-RU): Uses a different, potentially proprietary protocol (NOT Open Fronthaul)

Cooperative Transport(CTI)
    This handles a very specific problem..Sometimes the cable connecting DU and RU isnt a dedicated fiber. It might be a shared network like PON(Passive Optical Network) or a DOCSIS(cable tv) network, where bandwidth is shared with other people.
    
    O-RU needs to send upload data now, but shared network might be busy with someone else's traffic

    The Solution: The CTI (Cooperative Transport Interface).
    How it works:
        The O-DU knows exactly when the mobile phone is scheduled to send data.
        The O-DU sends a "heads up" message via the CTI to the Transport Node (TN) (e.g., the PON OLT).
        The TN reserves bandwidth in advance for that specific millisecond so the path is clear when the data arrives.

Inter O-DU Carrier Aggregation (The D2 Interface)
    Carrier Aggregation (CA) is a technique where a phone connects to two frequencies at once to double its speed (e.g., downloading on 2GHz and 3GHz simultaneously). Usually, these two frequencies are handled by the same O-DU.
    But what if the 2GHz radio is connected to O-DU #1 and the 3GHz radio is connected to O-DU #2?
    The Solution: The D2 Interface.
    Function: It creates a direct link between two O-DUs.
    Role: One O-DU acts as the PCell (Master) and the other as the SCell (Secondary). They coordinate user data directly so the phone can use both radios seamlessly.
    Condition: Both O-DUs must be connected to the same O-CU-CP.