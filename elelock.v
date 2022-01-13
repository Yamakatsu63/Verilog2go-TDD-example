module elelock(clk, key, close, lock);
    input clk, close;
    input [9:0] tenkey;
    reg key;
    output lock;

    always @(posedge clk) begin
        key <= keyenc(tenkey);
        if (key == 1'b1)
            lock <= 1'b0;
        else if (close == 1'b1)
            lock <= 1'b1;
    end

    function [1:0]keyenc;
    input [1:0] sw;
        case(sw)
            2'b01: keyenc = 1'b0;
            2'b10: keyenc = 1'b1;
        endcase
    endfunction
endmodule