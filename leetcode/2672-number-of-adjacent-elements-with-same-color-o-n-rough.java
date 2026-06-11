import java.util.Arrays;

class Solution {
	static int[] colorTheArray(int n, int[][] queries) {
		int[] colors = new int[n];
		for (int i=0;i<colors.length;i++) {
			colors[i] = -1*(i+1);

		}
		int[] adjacentElementsWithSameColor = new int[queries.length];
		if (n == 0) {
			return adjacentElementsWithSameColor; 
		}
		int adjacentElementsCount = 0;
		for(int i=0; i < queries.length; i++) {
			int index = queries[i][0];
			int color = queries[i][1];

			if (index == 0  && colors[0] == colors[1] && color != colors[1]) {
				colors[index] = color;
				adjacentElementsCount--; 
				continue;
			} else if (index == 0) {
				colors[index] = color;
				continue;
			} else if (index == colors.length-1 && colors[colors.length-1] == colors[colors.length-2] && color != colors[colors.length-2]) {
				colors[index] = color;
				adjacentElementsCount--; 
				continue;
			} else if (index == colors.length-1) {
				colors[index] = color;
				continue;
			}

			if (allSame(colors, index)) {
				colors[index] = color;
				if(!allSame(colors, index)) {
					adjacentElementsCount-=2;
				} 
			} else if (leftSideSame(colors, index) || rightSideSame(colors, index)) {
				colors[index] = color;
				if(allDifferent(colors, index)) {
					adjacentElementsCount-=1;
				}

			} else {
				colors[index] = color;
				if (leftSideSame(colors, index) || rightSideSame(colors, index)) {
					adjacentElementsCount+=1;
				}
			}
	         	adjacentElementsWithSameColor[i] = adjacentElementsCount;

		}
		return  adjacentElementsWithSameColor;

	}

	static boolean allSame(int[] colors, int queryIndex) {
		int left = queryIndex - 1;
		int right = queryIndex + 1;
		if (colors[left] == colors[queryIndex] && colors[queryIndex] == colors[right]) {
			return true;
		} else {
			return false;
		}
	}

	static boolean allDifferent(int[] colors, int queryIndex) {
		int left = queryIndex - 1;
		int right = queryIndex + 1;
		if (colors[left] != colors[queryIndex] && colors[queryIndex] == colors[right] && colors[left] != colors[right]) {
			return true;
		} else {
			return false;
		}
	}

	static boolean leftSideSame(int[] colors, int queryIndex) {
		int left = queryIndex - 1;
		if (colors[left] == colors[queryIndex]) {
			return true;
		} else {
			return false;
		}
	}

	static boolean rightSideSame(int[] colors, int queryIndex) {
		int right = queryIndex + 1;
		if (colors[right] == colors[queryIndex]) {
			return true;
		} else {
			return false;
		}
	}

	public static void main(String args[]) {
		Solution sol = new Solution();
		int[][] queries = {{0,2}, {1,2}, {3,1}, {1,1}, {2,1}};
		int[] answer = sol.colorTheArray(5, queries);
		System.out.println(Arrays.toString(answer));
	}
}

