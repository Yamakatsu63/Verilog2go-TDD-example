module elelock(key, lock);
    input key;
    output lock;

    assign lock = !key;
endmodule