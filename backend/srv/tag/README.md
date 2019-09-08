# TAG service
## 依赖安装

    pip install --user flask
    
其余缺少自行安装
    
## 数据集
  [chinese_L-12_H-768_A-12.tar.gz](https://jbox.sjtu.edu.cn/l/aoMMkf) 解压到`data`目录下， [sjtujj.tar.gz](https://jbox.sjtu.edu.cn/l/lu8AJY) 解压到`data/model`目录下

## 使用
- 启动服务： 运行`python serve.py`
- 使用：`http://localhost:5000/?text=xxx`
- 训练数据： 训练数据暂未上传
- 可选模型：
    - train/54.53 & test/55.58_40w_categorical
    - train/92.79 & test/85.99_40w_top3
    - train/97.85 & test/92.39_40w_top5
    - ***train/98.57 & test/97.23_40w_top5(当前使用模型)***
- 模型说明：

```
________________________________________________________________________________________________________________________
Layer (type)                           Output Shape               Param #       Connected to                            
========================================================================================================================
Input-Token (InputLayer)               (None, 20)                 0                                                     
________________________________________________________________________________________________________________________
Input-Segment (InputLayer)             (None, 20)                 0                                                     
________________________________________________________________________________________________________________________
Embedding-Token (TokenEmbedding)       [(None, 20, 768), (21128,  16226304      Input-Token[0][0]                       
________________________________________________________________________________________________________________________
Embedding-Segment (Embedding)          (None, 20, 768)            1536          Input-Segment[0][0]                     
________________________________________________________________________________________________________________________
Embedding-Token-Segment (Add)          (None, 20, 768)            0             Embedding-Token[0][0]                   
                                                                                Embedding-Segment[0][0]                 
________________________________________________________________________________________________________________________
Embedding-Position (PositionEmbedding) (None, 20, 768)            15360         Embedding-Token-Segment[0][0]           
________________________________________________________________________________________________________________________
Embedding-Dropout (Dropout)            (None, 20, 768)            0             Embedding-Position[0][0]                
________________________________________________________________________________________________________________________
Embedding-Norm (LayerNormalization)    (None, 20, 768)            1536          Embedding-Dropout[0][0]                 
________________________________________________________________________________________________________________________
Encoder-1-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Embedding-Norm[0][0]                    
________________________________________________________________________________________________________________________
Encoder-1-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-1-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-1-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Embedding-Norm[0][0]                    
                                                                                Encoder-1-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-1-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-1-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-1-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-1-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-1-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-1-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-1-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-1-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-1-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-1-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-1-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-2-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-1-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-2-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-2-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-2-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-1-FeedForward-Norm[0][0]        
                                                                                Encoder-2-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-2-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-2-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-2-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-2-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-2-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-2-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-2-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-2-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-2-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-2-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-2-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-3-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-2-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-3-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-3-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-3-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-2-FeedForward-Norm[0][0]        
                                                                                Encoder-3-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-3-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-3-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-3-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-3-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-3-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-3-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-3-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-3-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-3-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-3-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-3-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-4-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-3-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-4-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-4-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-4-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-3-FeedForward-Norm[0][0]        
                                                                                Encoder-4-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-4-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-4-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-4-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-4-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-4-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-4-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-4-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-4-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-4-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-4-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-4-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-5-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-4-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-5-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-5-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-5-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-4-FeedForward-Norm[0][0]        
                                                                                Encoder-5-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-5-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-5-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-5-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-5-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-5-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-5-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-5-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-5-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-5-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-5-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-5-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-6-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-5-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-6-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-6-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-6-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-5-FeedForward-Norm[0][0]        
                                                                                Encoder-6-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-6-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-6-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-6-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-6-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-6-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-6-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-6-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-6-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-6-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-6-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-6-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-7-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-6-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-7-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-7-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-7-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-6-FeedForward-Norm[0][0]        
                                                                                Encoder-7-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-7-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-7-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-7-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-7-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-7-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-7-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-7-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-7-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-7-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-7-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-7-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-8-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-7-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-8-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-8-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-8-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-7-FeedForward-Norm[0][0]        
                                                                                Encoder-8-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-8-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-8-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-8-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-8-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-8-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-8-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-8-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-8-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-8-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-8-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-8-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-9-MultiHeadSelfAttention (Mult (None, 20, 768)            2362368       Encoder-8-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-9-MultiHeadSelfAttention-Dropo (None, 20, 768)            0             Encoder-9-MultiHeadSelfAttention[0][0]  
________________________________________________________________________________________________________________________
Encoder-9-MultiHeadSelfAttention-Add ( (None, 20, 768)            0             Encoder-8-FeedForward-Norm[0][0]        
                                                                                Encoder-9-MultiHeadSelfAttention-Dropout
________________________________________________________________________________________________________________________
Encoder-9-MultiHeadSelfAttention-Norm  (None, 20, 768)            1536          Encoder-9-MultiHeadSelfAttention-Add[0][
________________________________________________________________________________________________________________________
Encoder-9-FeedForward (FeedForward)    (None, 20, 768)            4722432       Encoder-9-MultiHeadSelfAttention-Norm[0]
________________________________________________________________________________________________________________________
Encoder-9-FeedForward-Dropout (Dropout (None, 20, 768)            0             Encoder-9-FeedForward[0][0]             
________________________________________________________________________________________________________________________
Encoder-9-FeedForward-Add (Add)        (None, 20, 768)            0             Encoder-9-MultiHeadSelfAttention-Norm[0]
                                                                                Encoder-9-FeedForward-Dropout[0][0]     
________________________________________________________________________________________________________________________
Encoder-9-FeedForward-Norm (LayerNorma (None, 20, 768)            1536          Encoder-9-FeedForward-Add[0][0]         
________________________________________________________________________________________________________________________
Encoder-10-MultiHeadSelfAttention (Mul (None, 20, 768)            2362368       Encoder-9-FeedForward-Norm[0][0]        
________________________________________________________________________________________________________________________
Encoder-10-MultiHeadSelfAttention-Drop (None, 20, 768)            0             Encoder-10-MultiHeadSelfAttention[0][0] 
________________________________________________________________________________________________________________________
Encoder-10-MultiHeadSelfAttention-Add  (None, 20, 768)            0             Encoder-9-FeedForward-Norm[0][0]        
                                                                                Encoder-10-MultiHeadSelfAttention-Dropou
________________________________________________________________________________________________________________________
Encoder-10-MultiHeadSelfAttention-Norm (None, 20, 768)            1536          Encoder-10-MultiHeadSelfAttention-Add[0]
________________________________________________________________________________________________________________________
Encoder-10-FeedForward (FeedForward)   (None, 20, 768)            4722432       Encoder-10-MultiHeadSelfAttention-Norm[0
________________________________________________________________________________________________________________________
Encoder-10-FeedForward-Dropout (Dropou (None, 20, 768)            0             Encoder-10-FeedForward[0][0]            
________________________________________________________________________________________________________________________
Encoder-10-FeedForward-Add (Add)       (None, 20, 768)            0             Encoder-10-MultiHeadSelfAttention-Norm[0
                                                                                Encoder-10-FeedForward-Dropout[0][0]    
________________________________________________________________________________________________________________________
Encoder-10-FeedForward-Norm (LayerNorm (None, 20, 768)            1536          Encoder-10-FeedForward-Add[0][0]        
________________________________________________________________________________________________________________________
Encoder-11-MultiHeadSelfAttention (Mul (None, 20, 768)            2362368       Encoder-10-FeedForward-Norm[0][0]       
________________________________________________________________________________________________________________________
Encoder-11-MultiHeadSelfAttention-Drop (None, 20, 768)            0             Encoder-11-MultiHeadSelfAttention[0][0] 
________________________________________________________________________________________________________________________
Encoder-11-MultiHeadSelfAttention-Add  (None, 20, 768)            0             Encoder-10-FeedForward-Norm[0][0]       
                                                                                Encoder-11-MultiHeadSelfAttention-Dropou
________________________________________________________________________________________________________________________
Encoder-11-MultiHeadSelfAttention-Norm (None, 20, 768)            1536          Encoder-11-MultiHeadSelfAttention-Add[0]
________________________________________________________________________________________________________________________
Encoder-11-FeedForward (FeedForward)   (None, 20, 768)            4722432       Encoder-11-MultiHeadSelfAttention-Norm[0
________________________________________________________________________________________________________________________
Encoder-11-FeedForward-Dropout (Dropou (None, 20, 768)            0             Encoder-11-FeedForward[0][0]            
________________________________________________________________________________________________________________________
Encoder-11-FeedForward-Add (Add)       (None, 20, 768)            0             Encoder-11-MultiHeadSelfAttention-Norm[0
                                                                                Encoder-11-FeedForward-Dropout[0][0]    
________________________________________________________________________________________________________________________
Encoder-11-FeedForward-Norm (LayerNorm (None, 20, 768)            1536          Encoder-11-FeedForward-Add[0][0]        
________________________________________________________________________________________________________________________
Encoder-12-MultiHeadSelfAttention (Mul (None, 20, 768)            2362368       Encoder-11-FeedForward-Norm[0][0]       
________________________________________________________________________________________________________________________
Encoder-12-MultiHeadSelfAttention-Drop (None, 20, 768)            0             Encoder-12-MultiHeadSelfAttention[0][0] 
________________________________________________________________________________________________________________________
Encoder-12-MultiHeadSelfAttention-Add  (None, 20, 768)            0             Encoder-11-FeedForward-Norm[0][0]       
                                                                                Encoder-12-MultiHeadSelfAttention-Dropou
________________________________________________________________________________________________________________________
Encoder-12-MultiHeadSelfAttention-Norm (None, 20, 768)            1536          Encoder-12-MultiHeadSelfAttention-Add[0]
________________________________________________________________________________________________________________________
Encoder-12-FeedForward (FeedForward)   (None, 20, 768)            4722432       Encoder-12-MultiHeadSelfAttention-Norm[0
________________________________________________________________________________________________________________________
Encoder-12-FeedForward-Dropout (Dropou (None, 20, 768)            0             Encoder-12-FeedForward[0][0]            
________________________________________________________________________________________________________________________
Encoder-12-FeedForward-Add (Add)       (None, 20, 768)            0             Encoder-12-MultiHeadSelfAttention-Norm[0
                                                                                Encoder-12-FeedForward-Dropout[0][0]    
________________________________________________________________________________________________________________________
Encoder-12-FeedForward-Norm (LayerNorm (None, 20, 768)            1536          Encoder-12-FeedForward-Add[0][0]        
________________________________________________________________________________________________________________________
non_masking_layer_1 (NonMaskingLayer)  (None, 20, 768)            0             Encoder-12-FeedForward-Norm[0][0]       
________________________________________________________________________________________________________________________
reshape_1 (Reshape)                    (None, 20, 768, 1)         0             non_masking_layer_1[0][0]               
________________________________________________________________________________________________________________________
conv2d_1 (Conv2D)                      (None, 18, 1, 300)         691500        reshape_1[0][0]                         
________________________________________________________________________________________________________________________
conv2d_2 (Conv2D)                      (None, 17, 1, 300)         921900        reshape_1[0][0]                         
________________________________________________________________________________________________________________________
conv2d_3 (Conv2D)                      (None, 16, 1, 300)         1152300       reshape_1[0][0]                         
________________________________________________________________________________________________________________________
max_pooling2d_1 (MaxPooling2D)         (None, 1, 1, 300)          0             conv2d_1[0][0]                          
________________________________________________________________________________________________________________________
max_pooling2d_2 (MaxPooling2D)         (None, 1, 1, 300)          0             conv2d_2[0][0]                          
________________________________________________________________________________________________________________________
max_pooling2d_3 (MaxPooling2D)         (None, 1, 1, 300)          0             conv2d_3[0][0]                          
________________________________________________________________________________________________________________________
concatenate_1 (Concatenate)            (None, 1, 1, 900)          0             max_pooling2d_1[0][0]                   
                                                                                max_pooling2d_2[0][0]                   
                                                                                max_pooling2d_3[0][0]                   
________________________________________________________________________________________________________________________
flatten_1 (Flatten)                    (None, 900)                0             concatenate_1[0][0]                     
________________________________________________________________________________________________________________________
dropout_1 (Dropout)                    (None, 900)                0             flatten_1[0][0]                         
________________________________________________________________________________________________________________________
dense_1 (Dense)                        (None, 1175)               1058675       dropout_1[0][0]                         
========================================================================================================================
Total params: 105,123,575
Trainable params: 105,123,575
Non-trainable params: 0
```