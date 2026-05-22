import java.util.Arrays;

class Solution {
	static int[] colorTheArray(int n, int[][] queries) {
		int[] colors = new int[n];
		int[] adjacentElementsWithSameColor = new int[queries.length];
		for(int i=0; i < queries.length; i++) {
			int[] query = queries[i];
			colors[query[0]] = query[1];
			adjacentElementsWithSameColor[i] =  countAdjacentPairs(colors);
		}
		return  adjacentElementsWithSameColor;

	}

	static int countAdjacentPairs(int[] colors) {
		int currentColor = colors[0];
		int adjacentElementsWithSameColor = 0;
		for (int i=1;i<colors.length;i++) {
			if(colors[i] == currentColor && currentColor != 0) {
				adjacentElementsWithSameColor++;
			} else {
				currentColor = colors[i];
			}
		}
		return adjacentElementsWithSameColor;

	}
	public static void main(String args[]) {
		Solution sol = new Solution();
		int[][] queries = {{0,2}, {1,2}, {3,1}, {1,1}, {2,1}};
		int[] answer = sol.colorTheArray(5, queries);
		System.out.println(Arrays.toString(answer));
	}
}

