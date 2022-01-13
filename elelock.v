module elelock(clk, key, close, lock);
    input clk, close;
    input [1:0] key;
    output lock;

    always @(posedge clk) begin
        if (key[1] == 1'b1)
            lock <= 1'b0;
        else if (close == 1'b1)
            lock <= 1'b1;
    end
endmodule