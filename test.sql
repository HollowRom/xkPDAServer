use AIS20220623135859


select top 10 *
from xkPdaServer_good_tool
where FMATERIALID = 118405 --and FUSEORGID = 1 and FISINVENTORY = 1
select top 10 *
from t_BD_MaterialStock
where FMATERIALID = 118405
select top 10 *
from T_BD_UNIT
where FMASTERID = 100007
select top 10 a.FMASTERID            as FMATERIALID,    --物料内码
              a.FNUMBER,                                --物料编码
              a1.FNAME,                                 --物料名称
              a1.FSPECIFICATION,                        --物料规格
              a2.FERPCLSID,
              case
                  when a2.FERPCLSID = 1 then N'外购'
                  when a2.FERPCLSID = 2 then N'自制'
                  when a2.FERPCLSID = 3 then N'委外'
                  else N'未定义' end as FErpClass,      --物料属性
              a2.FSUITE,                                --套件
              a2.FBASEUNITID,--物料基本单位内码
              a3.FNUMBER             as FBaseUnitNumber,--物料基本单位
              a2.FWEIGHTUNITID,                         --物料重量单位
              a2.FVOLUMEUNITID,--物料尺寸单位
              a2.FISPURCHASE,--允许采购
              a2.FISINVENTORY,--允许库存
              a2.FISSUBCONTRACT,--允许委外
              a2.FISSALE,--允许销售
              a2.FISPRODUCE,--允许生产
              a2.FISASSET,--允许资产
              a4.FISBATCHMANAGE,                        --是否启用批号管理
              a4.FISKFPERIOD,                           --是否启用保质期管理
              a5.FCHECKINCOMING,--来料检验
              a5.FCHECKPRODUCT,--产品检验,
              a5.FCHECKSTOCK,--库存检验,
              a5.FCHECKRETURN,--退货检验
              a5.FCHECKDELIVERY,--发货检验
              a41.FSTOCKID,                             --默认仓库
              a41.FNAME              as FStockName,     --默认仓库名称
              a.FUSEORGID,                              --使用组织内码
              b.FNUMBER              as FUseOrgNumber,  --使用组织编码
              c.FNAME                as FUseOrgName     --使用组织名称
from T_BD_MATERIAL a,
     T_BD_MATERIAL_L a1,
     t_BD_MaterialBase a2,
     T_BD_UNIT a3
         left join T_BD_STOCK_L a41 on a3.FMASTERID = a41.FSTOCKID and a41.FLOCALEID = 2052,
     t_BD_MaterialStock a4,
     T_BD_MATERIALQUALITY a5,
     T_ORG_organizations b,
     T_ORG_organizations_L c
where a.FUSEORGID = b.FORGID
  and a.FDOCUMENTSTATUS = 'C'
  and a.FFORBIDSTATUS = 'A'
  and a.FMATERIALID = a1.FMATERIALID
  and a.FMATERIALID = a4.FMATERIALID
  and a.FMATERIALID = a2.FMATERIALID
  and a.FMATERIALID = a5.FMATERIALID
  and a2.FUSEORGID = a.FUSEORGID
  and a2.FBASEUNITID = a3.FMASTERID
  and a3.FUSEORGID = a.FUSEORGID
  and a1.FLOCALEID = 2052
  and a1.FUSEORGID = a.FUSEORGID
  and b.FORGID = c.FORGID
  and c.FLOCALEID = 2052


select top 10 *
from T_BD_MATERIAL
where FMASTERID = 118405

select top 10 a.FMASTERID as FMATERIALID, a3.FUSEORGID, a.FUSEORGID
from T_BD_MATERIAL a,
     T_BD_MATERIAL_L a1,
     t_BD_MaterialBase a2,
     T_BD_UNIT a3
         left join T_BD_STOCK_L a41 on a3.FMASTERID = a41.FSTOCKID and a41.FLOCALEID = 2052,
     t_BD_MaterialStock a4,
     T_BD_MATERIALQUALITY a5,
     T_ORG_organizations b,
     T_ORG_organizations_L c
where a.FUSEORGID = b.FORGID
  and a.FDOCUMENTSTATUS = 'C'
  and a.FFORBIDSTATUS = 'A'
  and a.FMATERIALID = a1.FMATERIALID
  and a.FMATERIALID = a4.FMATERIALID
  and a.FMATERIALID = a2.FMATERIALID
  and a.FMATERIALID = a5.FMATERIALID
  and a2.FUSEORGID = a.FUSEORGID
  and a2.FBASEUNITID = a3.FMASTERID
  and a3.FUSEORGID in (a.FUSEORGID, 0)
  and a1.FLOCALEID = 2052
  and a1.FUSEORGID = a.FUSEORGID
  and b.FORGID = c.FORGID
  and c.FLOCALEID = 2052
  and a.FMASTERID = 118405


select top 10 *
from T_BD_UNIT
where FNUMBER = 'zhi'--FUSEORGID = 0
select top 100 *
from xkPdaServer_userInfo_tool
select top 100 *
from xkPdaServer_keeperInfo_tool
select top 100 *
from xkPdaServer_empInfo_tool
select top 100 *
from xkPdaServer_customer_tool
select top 100 *
from xkPdaServer_sale_tz_to_stockout_tool
select top 10 *
from xkPdaServer_sub_ppbom_to_stockout_tool
select *
from xkPdaServer_poorder_to_stockin_tool

select top 10000 a.FSTOCKORGID         as 仓库组织,
                 a.FKEEPERID           as 保管组织,
                 b.FNUMBER             as 仓库编码,
                 c.FNAME               as 仓库名称,
                 isnull(d.FNUMBER, '') as 批号编码,
                 e.FMATERIALID         as 编码内码,
                 e.FNUMBER             as 编码编号,
                 f.FNAME               as 编码名称,
                 g.FNUMBER             as 单位编码,
                 a.FUPDATETIME         as 最后入库时间
from T_STK_INVENTORY a
         left join
     T_BD_LOTMASTER d on a.FLOT = d.FLOTID,
     T_BD_STOCK b,
     T_BD_STOCK_L c,
     T_BD_MATERIAL e,
     T_BD_MATERIAL_L f,
     T_BD_UNIT g
where a.FSTOCKID = b.FSTOCKID
  and b.FSTOCKID = c.FSTOCKID
  and c.FLOCALEID = 2052
  and a.FMATERIALID = e.FMATERIALID
  and e.FMATERIALID = f.FMATERIALID
  and f.FLOCALEID = 2052
  and a.FSTOCKSTATUSID = '10000'
  and a.FBASEUNITID = g.FMASTERID
  and a.FBASEQTY <> 0


select top 10 *
from T_STK_INVENTORY


select top 10 *
from xkPdaServer_gxplan_to_gxreport_tool a

select top 10 *
from T_SFC_OPERPLANNING


select top 10 a.FDEPARTMENTID as 加工车间内码, b.fnumber, *
from T_SFC_OPERPLANNINGDETAIL a, xkPdaServer_depart_tool b where a.FDEPARTMENTID = b.FMASTERID

select a.FMASTERID, a.FNUMBER, b.FNAME, c.FNUMBER as FUseOrgNumber
from T_BD_DEPARTMENT a,
     T_BD_DEPARTMENT_L b,
     T_ORG_organizations c
where a.FMASTERID = b.FDEPTID
  and b.FLOCALEID = 2052
  and a.FDOCUMENTSTATUS = 'C'
and a.FUSEORGID = c.FORGID
select * from    xkPdaServer_orgInfo_tool

select * from xkPdaServer_gxplan_to_gxreport_tool


select * from xkPdaServer_gxplan_to_gxreport_tool where FBILLNO = 'OP000771'


select t0.FID,       t0.FBILLNO,       t0.FMOID,       t0.FMOENTRYSEQ,       t0.FMOENTRYID,       t1.FENTRYID,       t1.FSEQ,       t2.FDEPARTMENTID, --加工车间
       t1.FSEQNUMBER,--工作序列
       t2.FOPERUNITID,--工序单位
       t2.FWORKCENTERID, --工作中心
       t2.FOPERNUMBER,   --工序序号
       t3.FMATERIALID,       t3.FNUMBER,       t3.FNAME,       t3.FSPECIFICATION,       t3.FBaseUnitNumber,       iif(t3.FISBATCHMANAGE = '1', t0.FLOT_TEXT, '')   as FLOT_TEXT,       cast(convert(float, t1.FSEQQTY) as varchar(24))  as FMustQty,
       cast(convert(float, t1.FSEQQTY) as varchar(24))  as SQTY,       t3.FUseOrgNumber,       t3.FISBATCHMANAGE,       row_number() over (order by t0.FMODIFYDATE desc) as idx
from T_SFC_OPERPLANNING t0
         left join T_SFC_OPERPLANNINGSEQ t1 on t0.FID = t1.FID
         left join T_SFC_OPERPLANNINGDETAIL t2 on t1.FENTRYID = t2.FENTRYID
         left join T_SFC_OPERPLANNINGDETAIL_A t2_A on t2.FDETAILID = t2_A.FDETAILID
         left join T_SFC_OPERPLANNINGDETAIL_D t2_D on t2.FDETAILID = t2_D.FDETAILID
         left join T_SFC_OPERPLANNINGDETAIL_b t2_B on t2.FENTRYID = t2_B.FENTRYID
         join xkPdaServer_good_tool t3 on t3.FMATERIALID = t0.FProductId and t0.FPROORGID in (t3.FUSEORGID, 0)
where /*isnull(t2.FPROCESSORGID, 0) in (0, 1)
  and*/
    /*t2.FOPTCTRLCODEID in (select FId from T_ENG_OPTCTRLCODE where FReportMode <> '40')
  and*/ t0.FMOENTRYID in (select FEntryID from T_PRD_MOENTRY_A where FStatus in (3, 4, 5))
  and t2.FWorKCENTERID in (select DSP1.FWCID
                           from T_SFC_DSPRPTPERMENTRY DSP1
                                    inner join T_SFC_DSPRPTPERM DSP0 on DSP1.FID = DSP0.FID
                           where DSP0.FUSERID = 102112
                             and DSP1.FISCHECKED = '1')
  and t0.FBILLTYPE in (' ', '001f29d2c9af844211e342cad266ac71')
 and t2_D.FISDisCRETEOPERDisPDETAIL = '0'
  and t0.FDOCUMENTSTATUS = N'C'
  and t2.FOPERCANCEL = N'A'
  and t0.FMOisSUSPEND = N'0'
  and t2_A.FISOUTSRC = N'0'
  and t2.FOPERSTATUS in ('3', '4', '5')
  and t0.FForMID = 'SFC_OperationPlaNNING'
  and t2_b.FREPORTQTY < t0.FMOQTY
  and t0.FBILLNO = 'OP000771'

select t0.FProductId,*
from T_SFC_OPERPLANNING t0
         left join T_SFC_OPERPLANNINGSEQ t1 on t0.FID = t1.FID
         left join T_SFC_OPERPLANNINGDETAIL t2 on t1.FENTRYID = t2.FENTRYID
         left join T_SFC_OPERPLANNINGDETAIL_A t2_A on t2.FDETAILID = t2_A.FDETAILID
         left join T_SFC_OPERPLANNINGDETAIL_D t2_D on t2.FDETAILID = t2_D.FDETAILID
         left join T_SFC_OPERPLANNINGDETAIL_b t2_B on t2.FDETAILID = t2_B.FDETAILID
         join xkPdaServer_good_tool t3 on t3.FMATERIALID = t0.FProductId --and t0.FPROORGID in (t3.FUSEORGID, 0)
where t0.FBILLNO = 'OP000771'

select * from dbo.T_BD_MATERIAL where fnumber = '1.01.001'

update T_BD_MATERIAL set T_BD_MATERIAL.FDOCUMENTSTATUS = 'C' where fnumber = '1.01.001'
