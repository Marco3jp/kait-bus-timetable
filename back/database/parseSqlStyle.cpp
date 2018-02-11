#include <bits/stdc++.h>
#define ll long long
#define all(a) (a).begin(),(a).end()

using namespace std;

int main(int argc, char const *argv[]) {
    if(argc != 3 || argc != 4){
        printf("引数の数がおかしいです。よく確認してください。\n");
        return 0;
    }

    std::vector<int> timeTable;
    int tmpH,tmpM;
    while (1) {
        scanf("%d:%d", &tmpH,&tmpM);
        if(tmpH == -1 && tmpM == -1){
            break;
        }
        timeTable.push_back(tmpH*60+tmpM);
    }

    if(argc == 3){
        for (size_t i = 0; i < timeTable.size(); i++) {
            printf("(%d,%d,%d),", timeTable[i] ,atoi(argv[1]),atoi(argv[2])); //帰り
        }
    }else if(argc == 4){
        for (size_t i = 0; i < timeTable.size(); i++) {
            printf("(%d,%d,%d,%d),", timeTable[i] ,atoi(argv[1]),atoi(argv[2]),atoi(argv[3])); //行き
        }
    }

    printf("\n");
    return 0;
}

//行き : time,from(id),to(id),daytype(id)
//帰り : time,from(id),daytype
//引数はfromとtoとdaytypeを必要なだけ入れてください。順番だけは守ってください。
