using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Collections;
using System.Collections.Generic;

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
class Solution
{
    static void Main(string[] args)
    {
        int L = int.Parse(Console.ReadLine());
        int H = int.Parse(Console.ReadLine());
        
        string T = Console.ReadLine();
        T = T.ToUpper();

        var index = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?";
        
        for (int i = 0; i < H; i++)
        {
            string ROW = Console.ReadLine();
            
            for (int j =0; j < T.Length; j++)
            {
                var ch = T[j];
                
                if (ch < 'A' || ch > 'Z')
                {
                    ch = '?';
                }
                
                Console.Write(ROW.Substring(L * index.IndexOf(ch), L));
            }
            
            Console.WriteLine();
        }
    }
}
