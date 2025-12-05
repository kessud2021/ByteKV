import React, { useEffect, useState } from "react";

export default function App() {
    const [keys, setKeys] = useState([]);
    const [selectedKey, setSelectedKey] = useState(null);
    const [value, setValue] = useState("");
    const [search, setSearch] = useState("");

    const api = {
        list: async () => {
            let r = await fetch("http://localhost:8080/keys");
            setKeys(await r.json());
        },
        get: async (k) => {
            let r = await fetch("http://localhost:8080/get?key=" + k);
            let v = await r.text();

            setSelectedKey(k);
            setValue(v);
        },
        set: async (k, v) => {
            await fetch("http://localhost:8080/set", {
                method: "POST",
                body: JSON.stringify({ key: k, value: v }),
            });
            api.list();
        },
        del: async (k) => {
            await fetch("http://localhost:8080/delete?key=" + k, {
                method: "DELETE",
            });
            setSelectedKey(null);
            api.list();
        },
    };

    useEffect(() => {
        api.list();
    }, []);

    return (
        <div className="app-container">

            {/* SIDEBAR */}
            <div className="sidebar">
                <h1>ByteKV</h1>

                <div className="search-box">
                    <input
                        placeholder="Search keys..."
                        onChange={(e) => setSearch(e.target.value)}
                    />
                </div>

                <button
                    className="btn btn-new"
                    onClick={() => {
                        const k = prompt("Key name?");
                        const v = prompt("Value?");
                        if (k) api.set(k, v || "");
                    }}
                >
                    + New Key
                </button>

                <div className="key-list">
                    {keys
                        .filter((k) => k.includes(search))
                        .map((k) => (
                            <div
                                key={k}
                                className={`key-item ${
                                    selectedKey === k ? "active" : ""
                                }`}
                                onClick={() => api.get(k)}
                            >
                                {k}
                            </div>
                        ))}
                </div>
            </div>

            {/* MAIN PANEL */}
            <div className="main-panel">
                <h2>Dashboard</h2>

                {selectedKey ? (
                    <div>
                        <h3>{selectedKey}</h3>

                        <textarea
                            className="editor"
                            value={value}
                            onChange={(e) => setValue(e.target.value)}
                        />

                        <div style={{ marginTop: "15px", display: "flex", gap: "10px" }}>
                            <button
                                className="btn btn-save"
                                onClick={() => api.set(selectedKey, value)}
                            >
                                Save
                            </button>

                            <button
                                className="btn btn-delete"
                                onClick={() => api.del(selectedKey)}
                            >
                                Delete
                            </button>
                        </div>
                    </div>
                ) : (
                    <p>Select a key to view/edit.</p>
                )}

                <h3 style={{ marginTop: "40px" }}>Server Logs</h3>
                <div className="logs-box">Logs coming soon…</div>
            </div>
        </div>
    );
}
