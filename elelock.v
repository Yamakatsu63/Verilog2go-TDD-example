module elelock(key, close, lock);
    input key, close;
    output lock;

    assign lock = close;
endmodule