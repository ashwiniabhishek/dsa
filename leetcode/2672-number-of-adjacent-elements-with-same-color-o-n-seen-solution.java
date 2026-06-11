import java.util.*;

class Solution {
    public int[] colorTheArray(int n, int[][] queries) {
		int[] colors = new int[n];
		int count = 0;
		int[] counts = new int[queries.length];
		int i = 0;
		for (int[] query: queries) {
			int idx = query[0];
			int color = query[1];

			int prev = idx - 1 >= 0 ? idx - 1 : -1;
			int next = idx + 1 < n ? idx + 1 : n;

			if (prev != -1 && colors[idx] != 0  && colors[prev] == colors[idx] ) count--;
			if (next != n && colors[idx] != 0 && colors[next] == colors[idx]) count--;

			colors[idx] = color;

			if (prev != -1 && colors[idx] != 0  && colors[prev] == colors[idx] ) count++;
			if (next != n && colors[idx] != 0 && colors[next] == colors[idx]) count++;
			counts[i++] = count;
		}
		return counts;
        
    }

	public static void main(String[] args) {
		Solution s = new Solution();
		System.out.println(Arrays.toString(s.colorTheArray(4, new int[][]{{0,2}, {1, 2}, {3, 1}, {1, 1}, {2, 1}})));
		System.out.println(Arrays.toString(s.colorTheArray(1, new int[][]{{0,100000}})));
	}
}
