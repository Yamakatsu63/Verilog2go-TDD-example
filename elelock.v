module elelock(clk, key, close, lock);
    input clk, key, close;
    output lock;

    always @(posedge clk) begin
        if (key == 1'b1)
            lock <= 1'b0;
        else if (close == 1'b1)
            lock <= 1'b1;
    end
endmodule