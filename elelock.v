module elelock(clk, reset, key, close, lock);
    input clk, reset, close;
    input [9:0] tenkey;
    reg [3:0] key [0:3];
    output lock;
    wire match;

    parameter SECRET_0 = 4'h3, SECRET_1 = 4'h7;
    assign match = key[0] == SECRET_0 && key[1] == SECRET_1;

    always @(posedge clk or posedge reset) begin
        if (close == 1'b1 || reset == 1'b1)begin
            key[1] <= 4'b1111;
            key[0] <= 4'b1111;
        end
        else begin
            key[1] <= key[0];
            key[0] <= keyenc(tenkey);
        end
    end

    always @(posedge clk or posedge reset) begin
        if (match)
            lock <= 1'b0;
        else if (close == 1'b1 || reset == 1'b1)
            lock <= 1'b1;
    end

    function [3:0] keyenc;
    input [9:0] sw;
        case(sw)
            10'b0000000001: keyenc = 4'h0;
            10'b0000000010: keyenc = 4'h1;
            10'b0000000100: keyenc = 4'h2;
            10'b0000001000: keyenc = 4'h3;
            10'b0000010000: keyenc = 4'h4;
            10'b0000100000: keyenc = 4'h5;
            10'b0001000000: keyenc = 4'h6;
            10'b0010000000: keyenc = 4'h7;
            10'b0100000000: keyenc = 4'h8;
            10'b1000000000: keyenc = 4'h9;
        endcase
    endfunction
endmodule